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
                sh 'BRANCH=$(echo $GIT_BRANCH | cut -b 8-11)'
                sh 'echo $BRANCH'
            }
        }
        stage('Deploy') {
            steps {
                sh 'echo "Deploying..."'
            }
        }
    }
}