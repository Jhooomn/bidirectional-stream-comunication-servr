pipeline {
    agent any

    environment {
        GO_VERSION = "1.22.1"
        GOROOT = "/usr/local/go"
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.GOROOT}/bin:${env.GOPATH}/bin:${env.PATH}"
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
                    // Download and setup Go
                    sh '''
                        wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                        tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
                    '''
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
