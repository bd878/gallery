const methodsWithBody = [
  'POST', 'PUT', 'DELETE', 'PATCH',
]

function getFullUrl(url, isFullUrl) {
  // TODO: replace backend url with value from config
  return isFullUrl ? url : `http://galleryexample.com${url}`;
}

function prepareBody(body, method) {
  if (!methodsWithBody.includes(method)) {
    return;
  }

  if (body instanceof URLSearchParams) {
    return body;
  }

  return JSON.stringify({ ...body });
}

function getOptions(props) {
  const {
    headers,
    body,
    method,
    ...otherOptions
  } = props;

  return {
    headers: new Headers({ ...headers }),
    body: prepareBody(body, method),
    mode: 'no-cors',
    method,
    ...otherOptions
  };
}

export default function api(url, props = {}) {
  const { isFullUrl = false } = props;

  const fullUrl = getFullUrl(url, isFullUrl);
  const options = getOptions(props);

  return fetch(fullUrl, options)
    .then(res => {
      return res
        .text()
        .then(text => {
          try {
            return JSON.parse(text);
          } catch (e) {
            return text;
          }
        })
    })
    .catch(e => Promise.reject({
      error: e,
      message: 'Cannot parse the response',
    }));
}
