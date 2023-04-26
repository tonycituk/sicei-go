pipeline {
    agent any
    stages {
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
        stage('Build, Package & Tag') {
            steps {
                sh '''BRANCH=$(echo $GIT_BRANCH | cut -b 8-11)
                docker build -t sicei-go-$BRANCH:0.0.$BUILD_ID .'''
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker stop sicei-go-container'
                sh 'docker rm sicei-go-container'
                sh 'docker run --detach --name sicei-go-container -p 80:80 sicei-go-main:0.0.$BUILD_ID'
            }
        }
    }
}