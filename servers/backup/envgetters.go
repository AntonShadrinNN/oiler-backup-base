package backup

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

type EnvGetter interface {
	GetEnvs() []corev1.EnvVar
}

type EnvGetterMerger struct {
	envegetters []EnvGetter
}

func NewEnvGetterMerger(envgetters []EnvGetter) EnvGetterMerger {
	return EnvGetterMerger{
		envegetters: envgetters,
	}
}

func (egm EnvGetterMerger) GetEnvs() []corev1.EnvVar {
	envs := []corev1.EnvVar{}
	for _, getter := range egm.envegetters {
		envs = append(envs, getter.GetEnvs()...)
	}

	return envs
}

type CommonEnvGetter struct {
	DbUri        string
	DbPort       string
	DbUser       string
	DbPass       string
	DbName       string
	S3Endpoint   string
	S3AccessKey  string
	S3SecretKey  string
	S3BucketName string
	CoreAddr     string
}

func (ceg CommonEnvGetter) GetEnvs() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "DB_HOST",
			Value: ceg.DbUri,
		},
		{
			Name:  "DB_PORT",
			Value: ceg.DbPort,
		},
		{
			Name:  "DB_USER",
			Value: ceg.DbUser,
		},
		{
			Name:  "DB_PASSWORD",
			Value: ceg.DbPass,
		},
		{
			Name:  "DB_NAME",
			Value: ceg.DbName,
		},
		{
			Name:  "S3_ENDPOINT",
			Value: ceg.S3Endpoint,
		},
		{
			Name:  "S3_ACCESS_KEY",
			Value: ceg.S3AccessKey,
		},
		{
			Name:  "S3_SECRET_KEY",
			Value: ceg.S3SecretKey,
		},
		{
			Name:  "S3_BUCKET_NAME",
			Value: ceg.S3BucketName,
		},
		{
			Name:  "CORE_ADDR",
			Value: ceg.CoreAddr,
		},
	}
}

type BackuperEnvGetter struct {
	MaxBackupCount int
}

func (beg BackuperEnvGetter) GetEnvs() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "MAX_BACKUP_COUNT",
			Value: fmt.Sprint(beg.MaxBackupCount),
		},
	}
}

type RestorerEnvGetter struct {
	BackupRevision string
}

func (reg RestorerEnvGetter) GetEnvs() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "BACKUP_REVISION",
			Value: reg.BackupRevision,
		},
	}
}
