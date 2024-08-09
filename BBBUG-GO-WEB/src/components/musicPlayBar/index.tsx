import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { Button, ConfigProvider, Drawer, Slider } from "antd";
import { BarsOutlined } from "@ant-design/icons";

interface MusicPlayBarProps {}

const MusicPlayBar: React.FunctionComponent<MusicPlayBarProps> = (props) => {
  const {} = props;

  const [loading, setLoading] = useState(false);
  const [open, setOpen] = useState(false);
  const onClose = () => {
    setOpen(false);
  };

  const onSliderChange = () => {};

  return (
    <>
      <div className={"MusicPlayBar-panel"}>
        <div className={"MusicPlayBar-content"}>
          <div className={"MusicPlayBar-content-left"}>
            <div className={"MusicPlayBar-content-left-img"}>
              <img
                src="https://p1.music.126.net/MOEpPrfT5O4L0DUyEDzVGQ==/109951169592101392.jpg?param=130y130"
                alt=""
              />
            </div>
            <div className={"MusicPlayBar-content-left-info"}>
              <div className={"MusicPlayBar-content-left-info-name"}>
                <span>歌名</span>
                <span>歌手名</span>
              </div>
              <div className="progress-bar">
                <Slider min={1} max={20} onChange={onSliderChange} />
              </div>
            </div>
            <div className={"MusicPlayBar-content-right"}>
              <Button
                icon={<BarsOutlined className={"order-list-icon"} />}
                ghost
                onClick={() => {
                  setOpen(true);
                }}
              />
            </div>
          </div>
        </div>
      </div>
      <Drawer
        title="Basic Drawer"
        placement="right"
        closable={false}
        onClose={onClose}
        open={open}
        loading={loading}
        afterOpenChange={(visible) => !visible && setLoading(true)}
      >
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
      </Drawer>
    </>
  );
};

export default MusicPlayBar;
