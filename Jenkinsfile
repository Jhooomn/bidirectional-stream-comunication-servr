pipeline {
    agent any

    environment {
        GO_VERSION = "1.22.1"
        GO_DIR = "${env.WORKSPACE}/go"
        GOROOT = "${GO_DIR}/go"
        GOPATH = "${GO_DIR}/workspace"
        PATH = "${GOROOT}/bin:${GOPATH}/bin:${env.PATH}"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/Jhooomn/bidirectional-stream-comunication-servr.git'
            }
        }

        stage('Setup Go') {
            steps {
                script {
                    // Determine OS and download appropriate Go binary
                    def PLATFORM = sh(script: 'uname -s', returnStdout: true).trim()
                    def ARCH = sh(script: 'uname -m', returnStdout: true).trim()

                    def BASE_URL = "https://golang.org/dl/"
                    def GO_TAR
                    if (PLATFORM == 'Darwin' && ARCH == 'x86_64') {
                        GO_TAR = "go${GO_VERSION}.darwin-amd64.tar.gz"
                    } else if (PLATFORM == 'Linux' && ARCH == 'x86_64') {
                        GO_TAR = "go${GO_VERSION}.linux-amd64.tar.gz"
                    } else if (PLATFORM == 'Linux' && ARCH == 'aarch64') {
                        GO_TAR = "go${GO_VERSION}.linux-arm64.tar.gz"
                    } else {
                        error "Unsupported platform: ${PLATFORM} or architecture: ${ARCH}"
                    }

                    // Create the directory for Go and download the appropriate tarball
                    sh """
                        mkdir -p ${GO_DIR}
                        curl -L -o ${GO_TAR} ${BASE_URL}${GO_TAR}
                        tar -C ${GO_DIR} -xzf ${GO_TAR}
                    """
                }
            }
        }

        stage('Install govulncheck') {
            steps {
                sh 'go install golang.org/x/vuln/cmd/govulncheck@latest'
            }
        }

        stage('Run govulncheck') {
            steps {
                sh 'govulncheck ./...'
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
    }
}
