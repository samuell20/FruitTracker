export const FetchWrapper = {
    get,
    post,
    put,
    delete: _delete
};

async function get(url:string) {
    const requestOptions = {
        method: 'GET'
    };
    const response = await fetch(url, requestOptions);
    return handleResponse(response);
}

async function post(url:string, body:object) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body)
    };
    const response = await fetch(url, requestOptions);
    return handleResponse(response);
}

async function put(url:string, body:object) {
    const requestOptions = {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body)
    };
    const response = await fetch(url, requestOptions);
    return handleResponse(response);
}

async function _delete(url:string) {
    const requestOptions = {
        method: 'DELETE'
    };
    const response = await fetch(url, requestOptions);
    return handleResponse(response);
}


async function handleResponse(response:Response) {
    const text = await response.text();
    const data = text && JSON.parse(text);
    if (!response.ok) {
        const error = (data && data.message) || response.statusText;
        return Promise.reject(error);
    }
    return data;
}