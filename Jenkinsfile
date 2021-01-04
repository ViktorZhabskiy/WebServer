pipeline {
  agent any
  environment {
    GIT_COMMIT_HASH = sh(script: 'git rev-parse HEAD', , returnStdout: true).trim()
    DATE = sh(script: "date '+%Y-%m-%d'", , returnStdout: true).trim()

    TAGGED_IMG = "${DATE}:${GIT_COMMIT_HASH}"
    LATEST = "${DATE}:latest"
  }

  stages {
    stage('first step') {
      steps {
        script {
          sh '''ls -l'''
        }
      }
    }
  }
}