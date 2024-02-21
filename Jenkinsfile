// Run on an agent where we want to use Go
node {
    // Ensure the desired Go version is installed on this agent,
    // using the name defined in the Global Tool Configuration
    def root = tool type: 'go', name: '1.22'

    // Export environment variables to pointing the Go installation;
    // the `PATH+X` syntax prepends an item to the existing `PATH`:
    // https://jenkins.io/doc/pipeline/steps/workflow-basic-steps/#withenv-set-environment-variables
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        // Output will be something like "go version go1.19 darwin/arm64"
        sh 'go version'
    }
}
