pipeline {
    agent any

    environment {
        GOPATH = "${env.WORKSPACE}/go"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/Jhooomn/bidirectional-stream-comunication-servr.git'
            }
        }

        stage('Setup Go Environment') {
            steps {
                script {
                    // Use the configured Go tool from Jenkins
                    def goInstallation = tool(name: 'go-1.22.1', type: 'go')
                    env.GOROOT = "${goInstallation}"
                    env.PATH = "${env.GOROOT}/bin:${GOPATH}/bin:${env.PATH}"
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
