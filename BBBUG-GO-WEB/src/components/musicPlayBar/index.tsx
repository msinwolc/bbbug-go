import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { ConfigProvider, Slider } from "antd";

interface MusicPlayBarProps {}

const MusicPlayBar: React.FunctionComponent<MusicPlayBarProps> = (props) => {
  const {} = props;

  const onSliderChange = () => {};

  return (
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
              <Slider
                min={1}
                max={20}
                thumbStyle={{
                  width: 20, // 自定义滑块宽度
                  height: 1, // 自定义滑块高度
                  borderRadius: "50%", // 如果希望滑块为圆形，可以设置边框半径为高度的一半
                }}
                onChange={onSliderChange}
              />
            </div>
          </div>
          <div className={"MusicPlayBar-content-right"}></div>
        </div>
      </div>
    </div>
  );
};

export default MusicPlayBar;
