import React, {lazy} from 'react';
import Tag from '../Tag';
import i18n from '../../i18n';
import {getFileDownloadUrl} from "../../api";

const List = lazy(() => import("../../components/List"));
const ListItem = lazy(() => import("../../components/ListItem"));

const MessagesList = ({
  css,
  liCss,
  messages,
  loading,
  error,
}) => {
  let content = <Tag></Tag>;
  if (error) {
    content = <Tag>{error}</Tag>
  } else {
    content = (
      <List el="ul" css={css}>
        {messages.map(message => (
          <Tag
            el="li"
            css={liCss}
            key={message.id}
          >
            <Tag
              el="a"
              href={getFileDownloadUrl(`/messages/v1/read_file?id=${message.fileid}`, false)}
              download={message.filename}
              target="_blank"
            >
              {message.filename}
            </Tag>

            <ListItem key={message.id}>{message.value}</ListItem>
          </Tag>
        ))}
      </List>
    )
  }

  return (
    <>
      {loading ? <Tag>{i18n("loading")}</Tag> : null}
      {content}
    </>
  )
}

export default MessagesList;
