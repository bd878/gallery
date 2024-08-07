import React, {lazy} from 'react';
import api from '../../api';
import Tag from '../../components/Tag';
import i18n from '../../i18n';

const Form = lazy(() => import("../../components/Form"));
const FormField = lazy(() => import("../../components/FormField"));
const Button = lazy(() => import("../../components/Button"));

const sendLoginRequest = async e => {
  e.preventDefault();

  let name = e.target.elements['name'].value;
  let password = e.target.elements['password'].value;
  if (!name) {console.error(i18n("name_required_err")); return;}
  if (!password) {console.error(i18n("pass_required_err")); return;}

  const response = await api("/users/v1/login", {
    method: "POST",
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    },
    body: new URLSearchParams({
      'name': name,
      'password': password,
    })
  });
  console.log(response);
  if (response.status == "ok") {
    setTimeout(() => {location.href = "/home"}, 1000)
  }
}

const LoginForm = props => {
  return (
    <Tag>
      <Form name="login-form" onSubmit={sendLoginRequest}>
        <FormField required name="name" type="text" />
        <FormField required name="password" type="password" />
        <Button type="submit" text={i18n("login")} />
      </Form>

      <Tag el="a" href="/register" target="_self">{i18n("register")}</Tag>
    </Tag>
  );
}

export default LoginForm;
