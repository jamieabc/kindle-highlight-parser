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
	testOutput = "穷查理宝典：查理·芒格智慧箴言录 (查理·芒格)\n" +
		"\t- 简化任务的最佳方法一般是先解决那些答案显而易见的大问题。\n" +
		"\n" +
		"\t- 要得到你想要的某样东西，最可靠的办法是让你自己配得起\n" +
		"\n" +
		"\t- 要得到你想要的某样东西，最可靠的办法是让你自己配得起它。\n" +
		"\n" +
		"\t- 如果你们的正确让其他有身份有地位的人觉得没面子，那么你们可能会引发别人极大的报复心理。\n\n" +
		"“错误”的行为+助推（套装共二册）（2017诺贝尔经济学奖理查德·泰勒获奖作品） (理查德·塞勒)\n" +
		"\t- 我们10年以后享受到的快乐，同我们今天能够享受的快乐相比，其对我们的吸引力极为微小”。\n" +
		"\n" +
		"\t- 贴现效用模型的基本理念是，对你来说，即时消费比未来的消费更具价值。\n" +
		"\n" +
		"\t- 马修从纽约向西看，无法分辨是中国远还是日本远。但是如果他到了东京，则会发现从东京到上海的距离比从纽约到芝加哥还远。\n\n"
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