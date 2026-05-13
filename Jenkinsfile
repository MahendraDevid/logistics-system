pipeline {
    agent any

    environment {
        DOCKER_HUB_USER = 'arpusauri'
        DOCKER_HUB_ID   = 'docker-hub-credentials'
        KUBE_CONFIG_ID  = 'kube-config'
    }

    stages {

        stage('Checkout Repository') {
            steps {
                echo "=== Mengambil kode terbaru ==="
                git branch: 'arya',
                    url: 'https://github.com/MahendraDevid/logistics-system.git'
            }
        }

        // ===================================================================
        // ── 1. WAREHOUSE SERVICE ───────────────────────────────────────────
        // ===================================================================

        stage('WMS - Unit Test') {
            steps {
                echo "=== Unit Test: Warehouse Service ==="
                dir('warehouse-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('WMS - Lint') {
            steps {
                dir('warehouse-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('WMS - Build Image') {
            steps {
                dir('warehouse-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                        bat "docker build -t %DOCKER_HUB_USER%/warehouse-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('WMS - Functional Test') {
            steps {
                dir('warehouse-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    sleep time: 15, unit: 'SECONDS'
                    bat 'set TEST_DATABASE_URL=host=localhost user=testuser password=testpass dbname=wms_test port=5433 sslmode=disable && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('warehouse-service') { bat 'docker-compose -f deployments/docker-compose.test.yml down -v' }
                }
            }
        }

        stage('WMS - Push Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                    bat "docker push %DOCKER_HUB_USER%/warehouse-service:latest"
                }
            }
        }

        // ===================================================================
        // ── 2. SETTLEMENT SERVICE ──────────────────────────────────────────
        // ===================================================================

        stage('Settlement - Unit Test') {
            steps {
                echo "=== Unit Test: Settlement Service ==="
                dir('settlement-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('Settlement - Lint') {
            steps {
                dir('settlement-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('Settlement - Build Image') {
            steps {
                dir('settlement-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                        bat "docker build -t %DOCKER_HUB_USER%/settlement-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('Settlement - Functional Test') {
            steps {
                dir('settlement-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    sleep time: 15, unit: 'SECONDS'
                    bat 'set TEST_DATABASE_URL=host=localhost user=testuser password=testpass dbname=settlement_test port=5434 sslmode=disable && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('settlement-service') { bat 'docker-compose -f deployments/docker-compose.test.yml down -v' }
                }
            }
        }

        stage('Settlement - Push Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                    bat "docker push %DOCKER_HUB_USER%/settlement-service:latest"
                }
            }
        }

        // ===================================================================
        // ── 3. PRICING SERVICE ─────────────────────────────────────────────
        // ===================================================================

        stage('Pricing - Unit Test') {
            steps {
                echo "=== Unit Test: Pricing Service ==="
                dir('pricing-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('Pricing - Lint') {
            steps {
                dir('pricing-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('Pricing - Build Image') {
            steps {
                dir('pricing-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                        bat "docker build -t %DOCKER_HUB_USER%/pricing-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('Pricing - Functional Test') {
            steps {
                dir('pricing-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    sleep time: 15, unit: 'SECONDS'
                    bat 'set TEST_DATABASE_URL=host=localhost user=testuser password=testpass dbname=pricing_test port=5435 sslmode=disable && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('pricing-service') { bat 'docker-compose -f deployments/docker-compose.test.yml down -v' }
                }
            }
        }

        stage('Pricing - Push Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                    bat "docker push %DOCKER_HUB_USER%/pricing-service:latest"
                }
            }
        }

        // ===================================================================
        // ── 4. E-POD SERVICE ───────────────────────────────────────────────
        // ===================================================================

        stage('e-POD - Unit Test') {
            steps {
                echo "=== Unit Test: e-POD Service ==="
                dir('epod-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('e-POD - Lint') {
            steps {
                dir('epod-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('e-POD - Build Image') {
            steps {
                dir('epod-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                        bat "docker build -t %DOCKER_HUB_USER%/epod-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('e-POD - Functional Test') {
            steps {
                dir('epod-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    sleep time: 15, unit: 'SECONDS'
                    bat 'set TEST_DATABASE_URL=testuser:testpass@tcp(localhost:3307)/epod_test && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('epod-service') { bat 'docker-compose -f deployments/docker-compose.test.yml down -v' }
                }
            }
        }

        stage('e-POD - Push Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                    bat "docker push %DOCKER_HUB_USER%/epod-service:latest"
                }
            }
        }

        // ===================================================================
        // ── 5. ORDER MANAGEMENT SERVICE ────────────────────────────────────
        // ===================================================================

        stage('OMS - Unit Test') {
            steps {
                echo "=== Unit Test: Order Management Service ==="
                dir('order-management-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('OMS - Lint') {
            steps {
                dir('order-management-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('OMS - Build Image') {
            steps {
                dir('order-management-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                        bat "docker build -t %DOCKER_HUB_USER%/order-management-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('OMS - Functional Test') {
            steps {
                dir('order-management-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    sleep time: 15, unit: 'SECONDS'
                    bat 'set TEST_DATABASE_URL=host=localhost user=testuser password=testpass dbname=oms_test port=5436 sslmode=disable && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('order-management-service') {
                        bat 'docker-compose -f deployments/docker-compose.test.yml down -v'
                    }
                }
            }
        }

        stage('OMS - Push Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat "docker login -u %DOCKER_HUB_USER% -p %DOCKER_PASS%"
                    bat "docker push %DOCKER_HUB_USER%/order-management-service:latest"
                }
            }
        }

        // ===================================================================
        // ── 6. DEPLOYMENT & VERIFICATION ───────────────────────────────────
        // ===================================================================

        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([file(credentialsId: "${KUBE_CONFIG_ID}", variable: 'KUBECONFIG_FILE')]) {
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f warehouse-service/deployments/kubernetes/"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f settlement-service/deployments/kubernetes/"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f pricing-service/deployments/kubernetes/"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f epod-service/deployments/kubernetes/"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f order-management-service/deployments/kubernetes/"
                }
            }
        }

        stage('Verify Deployment') {
            steps {
                withCredentials([file(credentialsId: "${KUBE_CONFIG_ID}", variable: 'KUBECONFIG_FILE')]) {
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% get pods"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% get svc"
                    
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% rollout status deployment/warehouse-service --timeout=60s"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% rollout status deployment/settlement-service --timeout=60s"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% rollout status deployment/pricing-service --timeout=60s"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% rollout status deployment/epod-service --timeout=60s"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% rollout status deployment/order-management-service --timeout=60s"
                }
            }
        }
    }

    post {
        success {
            echo "✅ Pipeline berhasil! Seluruh service sudah ter-deploy ke Kubernetes."
        }
        failure {
            echo "❌ Pipeline gagal! Periksa log build di Jenkins untuk melihat tahap mana yang bermasalah."
        }
    }
}