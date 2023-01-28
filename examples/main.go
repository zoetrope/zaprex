package main

import (
	"errors"
	"flag"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	klog.InitFlags(nil)
	_ = flag.CommandLine.Parse([]string{})
	flag.Set("v", "10")

	// zap
	zapopts := zap.Options{
		Development: false,
		Level:       zapcore.Level(-10),
	}
	zapopts.BindFlags(flag.CommandLine)
	zapLogger := zap.New(zap.UseFlagOptions(&zapopts))
	zapLogger.Enabled()

	ctrl.SetLogger(zapLogger)
	klog.SetLogger(zapLogger.WithName("client-go"))

	ctrl.Log.Info("ctrl log0")
	ctrl.Log.V(1).Info("ctrl info1", "hoge", "fuga")
	ctrl.Log.V(2).Info("ctrl info2")
	ctrl.Log.V(3).Info("ctrl info3")
	ctrl.Log.Error(errors.New("ctrl"), "ctrl error")

	klog.Info("klog info0")
	klog.V(1).Info("klog info1", "hoge", "fuga")
	klog.V(2).Info("klog info2")
	klog.V(3).Info("klog info3")
	klog.InfoS("klog infoS", "hoge", "fuga", "sample", "test")
	klog.Errorf("klog error %v", errors.New("klog"))

}
