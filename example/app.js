// echo request body
exports.handler = async event => {
    // request body is json encoded string
    const body = JSON.parse(event.body);

    return {
        statusCode: 200,
        // response body is json string
        body: JSON.stringify(body),
        headers: {
            Location: 'http://google.com'
        }
    };
}
