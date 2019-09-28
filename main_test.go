package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	programName = "kindle-highlight-parser"
	sampleFile  = "test/My Clippings.txt"
	testFile    = "My Clippings.txt"
	outputFile  = "kindle.txt"
)

const (
	testOutput = "1. 穷查理宝典：查理·芒格智慧箴言录 (查理·芒格)\n" +
		"\n" +
		"\t- 大脑使用过度简单的运算法则的倾向对认知产生的糟糕影响，\n" +
		"\n" +
		"\t- 好像物理学忽略了（1）天体物理学，因为它的实验不可能在物理实验室中进行，（\n" +
		"\n" +
		"2. “错误”的行为+助推（套装共二册）（2017诺贝尔经济学奖理查德·泰勒获奖作品） (理查德·塞勒)\n" +
		"\n" +
		"\t- 房价严重偏离历史水平时，不管是偏高还是偏低，这些信号都有某种预测价值。\n" +
		"\n" +
		"\t- 投资者应该避免在市场显示出过热迹象时将资金投入其中，但也不应该期待通过短线操作而获得巨大收益。比起判断泡沫何时会破裂，发现我们是否处于泡沫当中其实更容易。\n" +
		"\n" +
		"\t- 理性经济人只会因为真实的信号而改变投资想法，但普通人可能会对那些都不能算作信号的事件做出反应，\n" +
		"\n" +
		"\t- 如果政策制定者只是简单地相信价格永远合理，那么他们永远都不会认为有政策干预的必要。但是，一旦我们承认可能会出现泡沫，并且私营企业似乎也正在助长这一疯狂的趋势，决策者在某种程度上出手干预就是有道理的。\n" +
		"\n"
)

func setupTestFile() {
	sample, _ := ioutil.ReadFile(sampleFile)
	_ = ioutil.WriteFile(testFile, sample, 0644)
}

func teardownTestFile() {
	_ = os.Remove(testFile)
}

func readFile() string {
	handle, _ := os.Open(outputFile)
	defer handle.Close()

	bytes, _ := ioutil.ReadAll(handle)
	return string(bytes)
}

func TestReadFile(t *testing.T) {
	setupTestFile()
	defer teardownTestFile()

	cmd := exec.Command(programName)
	out, err := cmd.CombinedOutput()
	assert.Equal(t, nil, err, "wrong command execution")
	assert.Equal(t, "", string(out), "wrong output")

	assert.Equal(t, testOutput, readFile(), "wrong output")
}
