import * as React from "react";
import "./index.less";

interface TopBannerProps {
  offsetWidth: number;
}

const TopBanner: React.FunctionComponent<TopBannerProps> = (props) => {
  const { offsetWidth } = props;
  return (
    <>
      <div className={"top-banner"}>
        <div className="animated-banner">
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/39f5846ccf601e178afe37eef2e9759d38078d56.png@1c.webp"
              data-height="281"
              data-width="3000"
              height="196"
              width="2100"
              referrerPolicy="no-referrer"
              style={{
                height: "196px",
                width: "2100px",
                transform: "translate(0px, 0px) rotate(0deg) scale(1)",
                opacity: 1,
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/0a3bf7a4c7c3b8b5b63aae6adf57c5fabeeb0cff.png@1c.webp"
              data-height="178"
              data-width="335"
              height="106"
              width="201"
              referrerPolicy="no-referrer"
              style={{
                height: "107px",
                width: "201px",
                transform: `translate(78px, 6px) rotate(0deg) scale(1)`,
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/e737e375a9bf412f802feaec18ee52ad517b4db9.png@1c.webp"
              data-height="92"
              data-width="249"
              height="23"
              width="62"
              referrerPolicy="no-referrer"
              style={{
                height: "23px",
                width: "62px",
                transform: `translate(${
                  -15 + offsetWidth / 10
                }px, 32px) rotate(0deg) scale(1)`,
                filter: "blur(0.5px)",
                opacity: "0.7",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/9539f58accbf88eb810a45ff2cbefbf1c29840c3.png@1c.webp"
              data-height="199"
              data-width="944"
              height="159"
              width="755"
              referrerPolicy="no-referrer"
              style={{
                height: "160px",
                width: "755px",
                transform: "translate(799.444px, -32px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/b5514fc043485ce31d0c91a398aeb288261256dc.png@1c.webp"
              data-height="170"
              data-width="1141"
              height="170"
              width="1141"
              referrerPolicy="no-referrer"
              style={{
                height: "170px",
                width: "1141px",
                transform: "translate(849px, -10px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/9cc0c1c3b87ca0ea32b80c3b08611e794108348b.png@1c.webp"
              data-height="92"
              data-width="249"
              height="46"
              width="124"
              referrerPolicy="no-referrer"
              style={{
                height: "46px",
                width: "124.5px",
                transform: "translate(224px, 30px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/8512395c322494ef32b0fb40821e8e66fb291b53.png@1c.webp"
              data-height="232"
              data-width="943"
              height="162"
              width="660"
              referrerPolicy="no-referrer"
              style={{
                height: "162.4px",
                width: "660.1px",
                transform: "translate(-562.309px, 0px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/e737e375a9bf412f802feaec18ee52ad517b4db9.png@1c.webp"
              data-height="92"
              data-width="249"
              height="55"
              width="149"
              referrerPolicy="no-referrer"
              style={{
                height: "55.2px",
                width: "149.4px",
                transform: "translate(-663.125px, 24px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/6e9dd4d7f26b117773661aeed3a357641d519f2a.png@1c.webp"
              data-height="223"
              data-width="2093"
              height="156"
              width="1465"
              referrerPolicy="no-referrer"
              style={{
                height: "156.1px",
                width: "1465.1px",
                transform:
                  "translate(-1.21528px, 0.243056px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/f48b2d7535cae1ed97c7bb3c69d29e574cbe30c6.png@1c.webp"
              data-height="159"
              data-width="410"
              height="87"
              width="225"
              referrerPolicy="no-referrer"
              style={{
                height: "87.45px",
                width: "225.5px",
                transform: "translate(-111.91px, 27.5px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/4d9b71139bc84c4ddd1b37a906115c4b0c54af87.png@1c.webp"
              data-height="205"
              data-width="385"
              height="102"
              width="192"
              referrerPolicy="no-referrer"
              style={{
                height: "102.5px",
                width: "192.5px",
                transform: "translate(346.528px, 35px) rotate(0deg) scale(1)",
                opacity: "0",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/cd873292b7b3ae5603b546e1ecdf05fecaa955f3.png@1c.webp"
              data-height="281"
              data-width="727"
              height="168"
              width="436"
              referrerPolicy="no-referrer"
              style={{
                height: "168.6px",
                width: "436.2px",
                transform: "translate(-424.167px, 0px) rotate(0deg) scale(1)",
                opacity: "0",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/f3a02af183b808d236d0cc7997cbeb8958ef9bb0.png@1c.webp"
              data-height="198"
              data-width="593"
              height="108"
              width="326"
              referrerPolicy="no-referrer"
              style={{
                height: "108.9px",
                width: "326.15px",
                transform: "translate(-390.729px, 22px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/cd873292b7b3ae5603b546e1ecdf05fecaa955f3.png@1c.webp"
              data-height="281"
              data-width="727"
              height="224"
              width="581"
              referrerPolicy="no-referrer"
              style={{
                height: "224.8px",
                width: "581.6px",
                transform: "translate(791.667px, 0px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/3483ce9cad9fc3dd80f34cb164bcd5eeb1606332.png@1c.webp"
              data-height="201"
              data-width="637"
              height="140"
              width="445"
              referrerPolicy="no-referrer"
              style={{
                height: "140.7px",
                width: "445.9px",
                transform: "translate(-147.292px, 0px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
          <div className="layer">
            <img
              src="https://i0.hdslb.com/bfs/vc/988a546df1cdf5eeefa0a5c319bc6b4bfca7d42d.png@1c.webp"
              data-height="201"
              data-width="637"
              height="140"
              width="445"
              referrerPolicy="no-referrer"
              style={{
                height: "140.7px",
                width: "445.9px",
                transform: "translate(550.278px, 0px) rotate(0deg) scale(1)",
                opacity: "1",
              }}
            />
          </div>
        </div>
        <div className="header-banner__inner">
          <a href="//www.bilibili.com" className="inner-logo">
            <img
              className="logo-img"
              alt="B站 b站"
              width="162"
              height="78"
              src="//i0.hdslb.com/bfs/archive/c8fd97a40bf79f03e7b76cbc87236f612caef7b2.png"
            />
          </a>
        </div>
      </div>
    </>
  );
};

export default TopBanner;
