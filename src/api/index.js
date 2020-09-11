import axios from 'axios';
​
class API {
  constructor(url) {
    let service = axios.create({
      baseURL: url,
      headers: {'X-Token': ''}
    });
    service.interceptors.response.use(this.handleSuccess, this.handleError);
    this.service = service;
  }
​
  setAuthToken = (token) => {
    this.token = token
  }
​
  handleSuccess = (response) => {
    return response;
  }
​
  handleError = (error) => {
    return Promise.reject(error)
  }
​
  redirectTo = (document, path) => {
    document.location = path
  }
​
  get = (path) => {
    return this.service.get(path, {
      method: 'GET',
      responseType: 'json',
      headers: {'X-Token': this.token}
    })
    .then(
      (response) => { return response }
    )
  }
​
  patch = (path, payload) => {
    return this.service.request(path, {
      method: 'PATCH',
      responseType: 'json',
      data: payload,
      headers: {'X-Token': this.token}
    }).then((response) => { return response });
  }
​
  post = (path, payload) => {
    return this.service.request(path, {
      method: 'POST',
      responseType: 'json',
      data: payload,
      headers: {'X-Token': this.token}
    }).then((response) => { return response });
  }
​
  delete = (path, payload) => {
    return this.service.request(path, {
      method: 'DELETE',
      responseType: 'json',
      data: payload,
      headers: {'X-Token': this.token}
    }).then((response) => { return response });
  }
}
​
export default new API('http://localhost:8080');