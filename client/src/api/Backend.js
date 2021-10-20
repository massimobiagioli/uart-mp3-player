import axios from "axios"

const reset = async () => axios.post('/api/reset')

const readSDContent = async () => {
    return axios
        .get('/api/sd-content')
        .then(response => response.data)
        .catch(error => {
            throw error;
        });
}

export {
    reset,
    readSDContent
}