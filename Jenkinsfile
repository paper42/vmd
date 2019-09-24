pipeline {
    agent any
    tools {
        go 'go-1.13'
    }
	environment {
		GOPATH = "${WORKSPACE}/go"
        GOCACHE = "/tmp/gocache"
        GO111MODULE = 'on'
    }
    stages {
        stage('Prepare') {
            steps {
                sh 'rm -rf $GOPATH'
                sh 'mkdir -p $GOPATH/src/github.com/bobertlo/'
                sh 'ln -s $WORKSPACE $GOPATH/src/github.com/bobertlo/vmd'
            }
        }
		stage('Compile') {
            steps {
                sh 'go build ./...'
            }
        }
		stage('Test') {
   	         steps {
   	             sh 'go test ./...'
   	         }
   	     }
	}
}
