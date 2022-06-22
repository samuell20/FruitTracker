import { FetchWrapper } from 'Helpers';

export const userService = {
    getAll,
    getById,
    create,
    update,
    delete: _delete
};

const baseUrl = `http://fruittrackerapp.tk:80/api/users`;

function getAll() {
    return FetchWrapper.get(baseUrl);
}

function getById(id:Number) {
    return FetchWrapper.get(`${baseUrl}/${id}`);
}

function create(params:object) {
    return FetchWrapper.post(baseUrl, params);
}

function update(id:Number, params:object) {
    return FetchWrapper.put(`${baseUrl}/${id}`, params);
}

// prefixed with underscored because delete is a reserved word in javascript
function _delete(id:Number) {
    return FetchWrapper.delete(`${baseUrl}/${id}`);
}
