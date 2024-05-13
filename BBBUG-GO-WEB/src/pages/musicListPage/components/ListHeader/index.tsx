import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import Search from "antd/es/input/Search";

interface ListHeaderProps {
  onSearch: (value: any) => void;
}

const ListHeader: React.FunctionComponent<ListHeaderProps> = (props) => {
  const { onSearch } = props;

  return (
    <div className={"ListHeader-panel"}>
      <Search placeholder="input search text" onSearch={onSearch} enterButton />
    </div>
  );
};

export default ListHeader;
