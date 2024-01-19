import TopBanner from "./pages/top-banner";
import "./App.css";
import TopMenuBar from "./pages/top-menu";
import { useEffect, useRef, useState } from "react";

function App() {
  const [isFix, setIsFix] = useState(false);
  const [offsetWidth, setOffsetWidth] = useState(0);
  const offsetWidthRef = useRef(0);
  offsetWidthRef.current = offsetWidth;

  const mainPageRef = useRef(null);
  const bannerRef = useRef(null);

  /**
   * 判断页面滑动时顶部菜单是否浮动
   * @param e
   */
  const handleScroll = () => {
    const mainPage: any = mainPageRef.current;
    mainPage.scrollTop > 50 ? setIsFix(true) : setIsFix(false);
  };
  /**
   * 处理鼠标移动事件
   * @param e 事件对象
   */
  const handleMouseEnter = (e: any) => {
    console.log("enter", e);
    // setOffsetWidth()
  };
  /**
   * 处理鼠标移动事件
   * @param e 事件对象
   */
  const handleMouseLeave = (e: any) => {
    console.log("leave", e);
    setOffsetWidth(0);
  };
  /**
   * 处理鼠标移动事件
   * @param e 事件对象
   */
  const handleMouseMove = (e: any) => {
    console.log("move", e.movementX + offsetWidthRef.current);
    setOffsetWidth(e.movementX + offsetWidthRef.current);
  };

  // 监听滚动事件
  useEffect(() => {
    const mainPage: any = mainPageRef.current;
    const banner: any = bannerRef.current;
    mainPage &&
      mainPage.addEventListener("scroll", () => {
        handleScroll();
      });
    banner &&
      banner.addEventListener("mouseleave", (e: any) => {
        handleMouseLeave(e);
      });
    banner &&
      banner.addEventListener("mousemove", (e: any) => {
        handleMouseMove(e);
      });
    return () => {
      mainPage.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <div className={"mainPage"} ref={mainPageRef}>
      <div ref={bannerRef}>
        <TopMenuBar isFix={isFix} menuData={["menu1", "menu2"]} />
        <TopBanner offsetWidth={offsetWidth} />
        <audio src="http://127.0.0.1:4000/api/song/playurl?mid=274791" controls></audio>
      </div>
      <div style={{ height: "2000px", width: "100%" }}></div>
    </div>
  );
}

export default App;
