import * as React from "react";
import "./index.less";
import { Card, List } from "antd";
import ListHeader from "./components/ListHeader";
import MusicPlayBar from "../../components/musicPlayBar";

interface MusicListPageProps {}

const MusicListPage: React.FunctionComponent<MusicListPageProps> = (props) => {
  const {} = props;

  const musicList = [
    {
      songId: 1,
      songName: "music1",
      songSinger: "singer1",
      songPic: "https://music.163.com/song/media/outer/url?id=1379680.mp3",
    },
    {
      songId: 2,
      songName: "music2",
      songSinger: "singer1",
      songPic: "https://music.163.com/song/media/outer/url?id=1379680.mp3",
    },
    {
      songId: 3,
      songName: "music3",
      songSinger: "singer1",
      songPic: "https://music.163.com/song/media/outer/url?id=1379680.mp3",
    },
    {
      songId: 4,
      songName: "music4",
      songSinger: "singer1",
      songPic: "https://music.163.com/song/media/outer/url?id=1379680.mp3",
    },
    {
      songId: 4,
      songName: "music4",
      songSinger: "singer1",
      songPic: "https://music.163.com/song/media/outer/url?id=1379680.mp3",
    },
    {
      songId: 4,
      songName: "music4",
      songSinger: "singer1",
      songPic: "https://music.163.com/song/media/outer/url?id=1379680.mp3",
    },
  ];
  // that.global.apiUrl + "/api/song/playurl?mid=" + obj.song.mid

  /**
   * 搜索逻辑处理
   * @param value
   */
  const onSearch = (value: any) => {
    console.log(value);
  };

  return (
    <div className={"musicListPagePanel"}>
      <div className={"musicCardPanel"}>
        <Card
          title="我喜欢的"
          extra={<a href="#">更多</a>}
          style={{ width: 650 }}
        >
          <p>Card content</p>
          <p>Card content</p>
          <p>Card content</p>
        </Card>
        <Card
          title="新歌推荐"
          extra={<a href="#">更多</a>}
          style={{ width: 650, marginLeft: 5 }}
        >
          <p>Card content</p>
          <p>Card content</p>
          <p>Card content</p>
        </Card>
      </div>
      <div className={"musicSearchListPanel"}>
        <List
          header={<ListHeader onSearch={onSearch} />}
          pagination={{ position: "bottom", align: "center" }}
          bordered
          dataSource={musicList}
          renderItem={(item) => <List.Item>{item.songName}</List.Item>}
        />
      </div>
      <div className={"musicPlayBar"}>
        <div className="musicPlayBar-content">
          <MusicPlayBar />
        </div>
      </div>
    </div>
  );
};

export default MusicListPage;
