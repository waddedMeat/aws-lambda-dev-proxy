/**
 * The handler parses the request body and echos it back in the expected format
 * @param event
 * @returns {Promise<{headers: {Location: string}, body: string, statusCode: number}>}
 */
export const handler = async event => {
    // request body is json encoded string
    const body = JSON.parse(event.body);

    return {
        statusCode: 200,
        // response body is json string
        body: JSON.stringify(body),
        // the headers are here for example purpose only
        headers: {
            Location: 'http://google.com'
        }
    };
}
