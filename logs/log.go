package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"go_learn_web/configs"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

/** 参考来源:
1. logrus https://blog.csdn.net/wslyk606/article/details/81670713
2. file-rotatelogs https://blog.csdn.net/jkwanga/article/details/107186653
*/

type lineHook struct {
	Field  string
	Skip   int
	levels []log.Level
}
type MyFormatter struct {
}

func (hook lineHook) Fire(entry *log.Entry) error {
	entry.Data["appName"] = "go_learn_web"
	hook.Field = "source"
	entry.Data[hook.Field] = findCaller(hook.Skip)
	return nil
}

func (hook lineHook) Levels() []log.Level {
	return log.AllLevels
}

// findCaller 处理调用链，返回log日志出处
func findCaller(skip int) string {
	file := ""
	line := 0
	var pc uintptr
	// 遍历调用链最大索引为1层
	for i := 0; i < 11; i++ {
		file, line, pc = getCaller(skip + i)
		// log包和logrus一起过滤
		if !strings.HasPrefix(file, "logrus") && !strings.HasPrefix(file, "log") {
			break
		}
	}

	fullFnName := runtime.FuncForPC(pc)

	fnName := ""
	if fullFnName != nil {
		fnNameStr := fullFnName.Name()
		// 取得函数名
		parts := strings.Split(fnNameStr, ".")
		fnName = parts[len(parts)-1]
	}
	return fmt.Sprintf("%s:%d:%s()", file, line, fnName)
}

// getCaller 获取文件信息
func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}
	n := 0

	// 获取包名
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, pc
}

// LoggerToFile gin日志中间件
func LoggerToFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		// 处理请求
		context.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := context.Request.Method
		reqUri := context.Request.RequestURI
		statusCode := context.Writer.Status()
		clientIP := context.ClientIP()
		log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func getFilePath() (fileName string) {
	logFilePath := configs.LOG_FILE_PATH
	logFileName := configs.LOG_FILE_NAME
	// 日志文件
	fileName = path.Join(logFilePath, logFileName)
	return
}

func Init() {
	fileName := getFilePath()
	// 尝试创建日志文件
	_, _ = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)

	// 写入文件, 使用rotatelogs进行日志分割
	logContent, err := rotatelogs.New(
		fileName+".%Y%m%d%H",
		rotatelogs.WithLinkName(fileName),      // 生成软链指向最新的日志文件
		rotatelogs.WithMaxAge(6*time.Hour),     // clear 最小单位为 分钟
		rotatelogs.WithRotationTime(time.Hour), // rotate 最小为一分钟轮询
	)
	if err != nil {
		log.Printf("falied to create rotatelogs: %s", err)
		return
	}

	log.AddHook(&lineHook{})
	// 设置json格式
	log.SetFormatter(&log.JSONFormatter{})
	// 输出到日志文件中
	log.SetOutput(logContent)
	// 日志级别是warn以上
	log.SetLevel(log.InfoLevel)
}
