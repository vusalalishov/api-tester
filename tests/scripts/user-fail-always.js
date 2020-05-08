function failIt(declarations, response) {
    return {
        exitCode: response.errorCode ? 11 : 0,
        message: "Failed"
    }
}