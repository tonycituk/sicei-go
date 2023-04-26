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
                sh 'docker images'
            }
        }
    }
}