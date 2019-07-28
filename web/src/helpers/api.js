import axios from 'axios';

import { APIConfig } from '../config';

const api = axios.create(APIConfig);

export default api;
