package stringEncryptor

type EncryptorWorker interface {
	doJob(workerId int,Jobs <-chan Job, results chan <-string)
	EncryptStrings([]string)(*[]string,*string)
	SetWorkersCount(count int)
	GetWorkersCount() int
}
const workersCountEVName="ENCR_WRKCNT"