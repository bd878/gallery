import React, {useEffect, useState} from 'react';
import api from '../../api';
import i18n from '../../i18n';

const Auth = props => {
  const [authed, setAuthed] = useState(false)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function call() {
      let response = { valid: false };
      try {
        setLoading(true);
        response = await api("/users/v1/auth", {
          method: 'POST',
          credentials: 'include',
        });
      } catch(e) {
        console.error("error occured on authing:", e);
      } finally {
        setLoading(false);
      }

      if (response.expired) {
        setAuthed(false);
        setTimeout(() => {location.href = "/login"}, 1000)
      } else {
        setAuthed(true);
        console.log("welcome,", response.user.name)
      }
    }

    call();
  }, [setAuthed, setLoading])

  if (loading) {
    return (<>{i18n('auth_process')}</>)
  }

  return (
    <>{authed
      ? props.children
      : (props.fallback || i18n("not_authed"))
    }</>
  );
}

export default Auth;
