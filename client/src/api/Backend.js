import axios from "axios"

const reset = async () => axios.post('/api/reset')

export {
    reset
}