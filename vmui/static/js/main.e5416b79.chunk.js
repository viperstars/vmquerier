(this.webpackJsonpvmui=this.webpackJsonpvmui||[]).push([[0],{339:function(e,t,n){},347:function(e,t,n){},349:function(e,t,n){"use strict";n.r(t);var a=n(2),c=n(0),r=n.n(c),i=n(24),o=n.n(i),s=(n(339),n(16)),l=n(422),u=n(420),j=Object(c.createContext)({showInfoMessage:function(){}}),d=function(){return Object(c.useContext)(j)},b=function(e){var t=e.children,n=Object(c.useState)({}),r=Object(s.a)(n,2),i=r[0],o=r[1],d=Object(c.useState)(!1),b=Object(s.a)(d,2),O=b[0],h=b[1],f=Object(c.useState)(void 0),x=Object(s.a)(f,2),p=x[0],v=x[1];Object(c.useEffect)((function(){p&&(o({message:p,key:(new Date).getTime()}),h(!0))}),[p]);return Object(a.jsxs)(j.Provider,{value:{showInfoMessage:v},children:[Object(a.jsx)(l.a,{open:O,autoHideDuration:4e3,onClose:function(e,t){"clickaway"!==t&&(v(void 0),h(!1))},children:Object(a.jsx)(u.a,{children:i.message})},i.key),t]})},O=n(412),h=n(402),f=n(416),x=n(92),p=n(390),v=n(404),m=n(383),y=n(385),g=n(386),S=n(425),E=n(206),T=n(183),w=n.n(T),A=n(13),C=n(12),R=n(62),_=n.n(R),k=n(181),U=n.n(k);_.a.extend(U.a);var D=[{long:"days",short:"d",possible:"day"},{long:"weeks",short:"w",possible:"week"},{long:"months",short:"M",possible:"mon"},{long:"years",short:"y",possible:"year"},{long:"hours",short:"h",possible:"hour"},{long:"minutes",short:"m",possible:"min"},{long:"seconds",short:"s",possible:"sec"},{long:"milliseconds",short:"ms",possible:"millisecond"}],M=D.map((function(e){return e.short})),I=function(e,t){var n=(t||new Date).valueOf()/1e3,a=e.trim().split(" ").reduce((function(e,t){var n=function(e){var t=e.match(/\d+/g),n=e.match(/[a-zA-Z]+/g);if(n&&t&&M.includes(n[0]))return Object(A.a)({},n[0],t[0])}(t);return n?Object(C.a)(Object(C.a)({},e),n):Object(C.a)({},e)}),{}),c=_.a.duration(a).asSeconds();return{start:n-c,end:n,step:c/30}},N=function(e){return new Date(1e3*e)},P=function(e,t){t?window.localStorage.setItem(e,JSON.stringify({value:t})):H([e])},L=function(e){var t=window.localStorage.getItem(e);if(null!==t)try{var n;return null===(n=JSON.parse(t))||void 0===n?void 0:n.value}catch(a){return t}},H=function(e){return e.forEach((function(e){return window.localStorage.removeItem(e)}))},F=["BASIC_AUTH_DATA","BEARER_AUTH_DATA"],B={serverUrl:L("PREFERRED_URL")||function(){var e=window.location.href.match(/^http.+\/vmui/)||["https://"];return Object(s.a)(e,1)[0].replace("vmui","prometheus")}(),displayType:"chart",query:L("LAST_QUERY")||"\n",time:{duration:"1h",period:I("1h")},queryControls:{autoRefresh:!1,autocomplete:L("AUTOCOMPLETE")||!1}};function z(e,t){switch(t.type){case"SET_DISPLAY_TYPE":return Object(C.a)(Object(C.a)({},e),{},{displayType:t.payload});case"SET_SERVER":return Object(C.a)(Object(C.a)({},e),{},{serverUrl:t.payload});case"SET_QUERY":return Object(C.a)(Object(C.a)({},e),{},{query:t.payload});case"SET_DURATION":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{duration:t.payload,period:I(t.payload,N(e.time.period.end))})});case"SET_UNTIL":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{period:I(e.time.duration,t.payload)})});case"SET_PERIOD":var n=function(e){var t=_.a.duration(e.to.valueOf()-e.from.valueOf());return["d","h","m","s"].map((function(e){return{val:t.get(e),str:e}})).filter((function(e){return 0!==e.val})).map((function(e){return"".concat(e.val).concat(e.str)})).join(" ")}(t.payload);return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{autoRefresh:!1}),time:Object(C.a)(Object(C.a)({},e.time),{},{duration:n,period:I(n,t.payload.to)})});case"TOGGLE_AUTOREFRESH":return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{autoRefresh:!e.queryControls.autoRefresh})});case"TOGGLE_AUTOCOMPLETE":return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{autocomplete:!e.queryControls.autocomplete})});case"RUN_QUERY":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{period:I(e.time.duration,N(e.time.period.end))})});case"RUN_QUERY_TO_NOW":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{period:I(e.time.duration)})});default:throw new Error}}var Y=n(114),q=n.n(Y),Q=function(e){if(/^(\d+|\d*\.\d+)$/.test(e))return parseFloat(e);var t={true:!0,false:!1,null:null,undefined:void 0};return e in t?t[e]:decodeURI(e)},W=function(e){var t=window;if(t){var n=t.location.protocol+"//"+t.location.host+t.location.pathname+"?"+e;t.history.pushState({path:n},"",n)}},G=Object(c.createContext)({}),J=function(){return Object(c.useContext)(G).state},V=function(){return Object(c.useContext)(G).dispatch},Z=Object.entries(B).reduce((function(e,t){var n=Object(s.a)(t,2),a=n[0],c=n[1];return Object(C.a)(Object(C.a)({},e),{},Object(A.a)({},a,function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:window.location.search;return q.a.parse(t,{ignoreQueryPrefix:!0,decoder:Q})[e]}(a)||c))}),{}),$=function(e){var t=e.children,n=Object(c.useReducer)(z,Z),r=Object(s.a)(n,2),i=r[0],o=r[1];Object(c.useEffect)((function(){!function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:window.location.search,n=q.a.parse(t,{ignoreQueryPrefix:!0,decoder:Q}),a=q.a.stringify(Object(C.a)(Object(C.a)({},n),e),{encode:!1});W(a)}(i)}),[i]);var l=Object(c.useMemo)((function(){return{state:i,dispatch:o}}),[i,o]);return Object(a.jsx)(G.Provider,{value:l,children:t})},K=function(e){return Object(a.jsxs)(f.a,{position:"relative",display:"inline-flex",children:[Object(a.jsx)(m.a,Object(C.a)({variant:"determinate"},e)),Object(a.jsx)(f.a,{top:0,left:0,bottom:0,right:0,position:"absolute",display:"flex",alignItems:"center",justifyContent:"center",children:Object(a.jsx)(x.a,{variant:"caption",component:"div",children:"".concat(e.label,"s")})})]})},X=n(384),ee=Object(X.a)({colorizing:{color:"white"}}),te=function(){var e=ee(),t=V(),n=J().queryControls.autoRefresh,i=Object(c.useState)(5),o=Object(s.a)(i,2),l=o[0],u=o[1],j=Object(c.useState)(),d=Object(s.a)(j,2),b=d[0],O=d[1],h=r.a.useState(100),x=Object(s.a)(h,2),p=x[0],v=x[1];Object(c.useEffect)((function(){var e;return n&&(O((new Date).valueOf()),e=setInterval((function(){O((new Date).valueOf()),t({type:"RUN_QUERY_TO_NOW"})}),1e3*l)),function(){e&&clearInterval(e)}}),[l,n]),Object(c.useEffect)((function(){var e=setInterval((function(){if(n&&b){var e=((new Date).valueOf()-b)/1e3,t=Math.floor(e/l*100);v(t)}}),16);return function(){clearInterval(e)}}),[n,b,l]);var m=function(){u((function(e){switch(e){case 1:return 2;case 2:return 5;case 5:return 1;default:return 5}}))};return Object(a.jsxs)(f.a,{display:"flex",alignItems:"center",children:[Object(a.jsx)(y.a,{control:Object(a.jsx)(g.a,{size:"small",className:e.colorizing,checked:n,onChange:function(){t({type:"TOGGLE_AUTOREFRESH"})}}),label:"Auto-refresh"}),n&&Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)(K,{className:e.colorizing,label:l,value:p,onClick:function(){m()}}),Object(a.jsx)(S.a,{title:"Change delay refresh",children:Object(a.jsx)(f.a,{ml:1,children:Object(a.jsx)(E.a,{onClick:function(){m()},children:Object(a.jsx)(w.a,{style:{color:"white"}})})})})]})]})},ne=n(186),ae=n.n(ne),ce=n(184),re=n.n(ce),ie=n(185),oe=n.n(ie),se=n(388),le=n(429),ue=n(10),je=Object(ue.a)({root:{padding:6,color:"white","&.Mui-selected":{color:"white"}}})(se.a),de=function(){var e=J().displayType,t=V();return Object(a.jsxs)(le.a,{value:e,exclusive:!0,onChange:function(n,a){return t({type:"SET_DISPLAY_TYPE",payload:null!==a&&void 0!==a?a:e})},children:[Object(a.jsxs)(je,{value:"chart","aria-label":"display as chart",children:[Object(a.jsx)(re.a,{}),"\xa0Query Range as Chart"]}),Object(a.jsxs)(je,{value:"code","aria-label":"display as code",children:[Object(a.jsx)(oe.a,{}),"\xa0Instant Query as JSON"]}),Object(a.jsxs)(je,{value:"table","aria-label":"display as table",children:[Object(a.jsx)(ae.a,{}),"\xa0Instant Query as Table"]})]})},be=n(25),Oe=n(27),he=n(46),fe=n.n(he),xe=n(106),pe=(n(347),n(187)),ve=function(e){var t=e.xScale,n=e.height,r=Object(c.useRef)(null);return Object(c.useEffect)((function(){Object(Oe.o)(r.current).call(Object(Oe.a)(t))}),[t]),Object(a.jsx)("g",{ref:r,className:"x axis",transform:"translate(0, ".concat(n,")")})},me=n(418),ye=function(e){var t=Math.abs(e),n=".2~s";return t>0&&t<.001&&(n=".0e"),t>=.001&&t<1&&(n=".3~f"),Object(me.a)(n)(e)},ge=function(e){var t=e.yScale,n=e.label,r=Object(c.useRef)(null);return Object(c.useEffect)((function(){t&&Object(Oe.o)(r.current).call(Object(Oe.b)(t).tickFormat(ye))}),[t]),Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)("g",{className:"y axis",ref:r}),n&&Object(a.jsx)("text",{style:{fontSize:"0.6rem"},transform:"translate(0,-2)",children:n})]})},Se=function(e){var t=e.height,n=e.x;return Object(a.jsx)(a.Fragment,{children:n&&Object(a.jsx)("line",{x1:n,y1:"0",x2:n,y2:t,stroke:"black",strokeDasharray:"4"})})},Ee=function(e){var t=e.yScale,n=e.xScale,r=e.datesInChart,i=e.onInteraction,o=e.setSelection,l=Object(c.useRef)(null),u=Object(c.useState)(),j=Object(s.a)(u,2),d=j[0],b=j[1],O=Object(c.useState)(),h=Object(s.a)(O,2),f=h[0],x=h[1],p=Object(c.useState)(!1),v=Object(s.a)(p,2),m=v[0],y=v[1];function g(e){var t=e.selection;if(t){if(!e.sourceEvent)return;y(!0);var a=t.map((function(e){return n.invert(e)})),c=Object(s.a)(a,2),r=c[0],i=c[1];o(r,i),Object(Oe.o)(l.current).call(E.move,null)}else y(!1)}var S=function(e){e.sourceEvent&&b(void 0)},E=Object(c.useMemo)((function(){return Object(Oe.d)().extent([[0,0],[n.range()[1],t.range()[0]]]).on("end",g).on("start",S)}),[g,n,t]);return Object(c.useEffect)((function(){var e=Object(Oe.c)((function(e){return e})).center;Object(Oe.o)(l.current).on("touchmove mousemove",(function(t){var a=Object(Oe.i)(t);m||(!function(t){var a=n.invert(t),c=e(r,a,1);b(c)}(a[0]),x(a[1]))})).on("mouseout",(function(){m||b(void 0)}))}),[n,r,m]),Object(c.useEffect)((function(){i(d,f)}),[d,f,i]),Object(c.useEffect)((function(){E&&n&&Object(Oe.o)(l.current).call(E)}),[n,E]),Object(a.jsx)(a.Fragment,{children:Object(a.jsx)("g",{ref:l})})},Te=n(389),we=Object(X.a)((function(){return{wrapper:{maxWidth:"40vw"}}})),Ae=function(e){var t=e.data,n=e.time,c=we();return Object(a.jsxs)(f.a,{px:1,className:c.wrapper,children:[Object(a.jsx)(f.a,{fontStyle:"italic",mb:.5,children:Object(a.jsx)(x.a,{variant:"subtitle1",children:"".concat(null===n||void 0===n?void 0:n.toLocaleDateString()," ").concat(null===n||void 0===n?void 0:n.toLocaleTimeString())})}),Object(a.jsx)(f.a,{mb:.5,my:1,children:Object(a.jsx)(x.a,{variant:"subtitle2",children:"Value: ".concat(new Intl.NumberFormat(void 0,{maximumFractionDigits:10}).format(t.value))})}),Object(a.jsx)(f.a,{children:Object(a.jsx)(x.a,{variant:"body2",children:t.metrics.map((function(e){var t=e.key,n=e.value;return Object(a.jsxs)(f.a,{component:"span",mb:.25,display:"flex",flexDirection:"row",alignItems:"center",children:[Object(a.jsxs)("span",{children:[t,":\xa0"]}),Object(a.jsx)("span",{style:{fontWeight:"bold"},children:n})]},t)}))})})]})},Ce=function(e){var t=e.series,n=e.timePresets,r=e.height,i=e.color,o=e.categories,l=Object(c.useState)(window.innerWidth),u=Object(s.a)(l,2),j=u[0],d=u[1],b=V(),O=10,h=20,x=40,p=50,v=Object(c.useMemo)((function(){return j-p-h}),[j,p,h]),m=Object(c.useMemo)((function(){return r-O-x}),[O,x]),y=Object(c.useMemo)((function(){return Object(Oe.m)().domain([n.start,n.end].map(N)).range([0,v])}),[v,n]),g=Object(c.useState)(!1),S=Object(s.a)(g,2),E=S[0],T=S[1],w=Object(c.useState)(),A=Object(s.a)(w,2),C=A[0],R=A[1],_=Object(c.useMemo)((function(){var e,n=t.reduce((function(e,t){return[].concat(Object(be.a)(e),Object(be.a)(t.values))}),[]).map((function(e){return e.value})),a=null!==(e=Object(Oe.g)(n))&&void 0!==e?e:1,c=Object(Oe.h)(n)||0;return Object(Oe.k)().domain([c>0?0:c,a<0?0:a]).range([m,0]).nice()}),[t,m]),k=Object(c.useMemo)((function(){return Object(Oe.f)().x((function(e){return y(N(e.key))})).y((function(e){return _(e.value||0)}))}),[y,_]),U=function(e){return k(e.values)},D=Object(c.useCallback)(function(){var e=Object(xe.a)(fe.a.mark((function e(n,a){var c,r,i,o;return fe.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:"number"===typeof n?a&&t&&t[0]&&(c=_.invert(a),r=t.map((function(e){var t;return null===(t=e.values[n])||void 0===t?void 0:t.value})).reduce((function(e,t,n){var a=Math.abs(c-t);return a<e.delta&&(e={delta:a,index:n}),e}),{delta:1/0,index:0}),i=N(t[0].values[n].key),o=i.valueOf()<(y.domain()[1].valueOf()+y.domain()[0].valueOf())/2,R({date:i,xCoord:y(i),index:n,activeSeries:r.index,leftPart:o}),T(!0)):(T(!1),R(void 0));case 1:case"end":return e.stop()}}),e)})));return function(t,n){return e.apply(this,arguments)}}(),[y,_,t]),M=Object(c.useMemo)((function(){return(null===C||void 0===C?void 0:C.activeSeries)?{value:t[C.activeSeries].values[C.index].value,metrics:o.map((function(e){return{key:e.key,value:t[C.activeSeries].metric[e.key]}}))}:void 0}),[C,t]),I=Object(c.useRef)(null),P=Object(c.useMemo)((function(){return t&&t[0]?t[0].values.map((function(e){return e.key})).map(N):[]}),[t]),L=function(e,t){b({type:"SET_PERIOD",payload:{from:e,to:t}})};return Object(a.jsx)(pe.a,{bounds:!0,onResize:function(e){var t=e.bounds;return t&&d(null===t||void 0===t?void 0:t.width)},children:function(e){var n=e.measureRef;return Object(a.jsxs)("div",{ref:n,style:{width:"100%"},children:[I&&M&&Object(a.jsx)(Te.a,{disableScrollLock:!0,style:{pointerEvents:"none"},id:"chart-tooltip-popover",open:E,anchorEl:I.current,anchorOrigin:{vertical:"top",horizontal:(null===C||void 0===C?void 0:C.leftPart)?20:-20},transformOrigin:{vertical:"top",horizontal:(null===C||void 0===C?void 0:C.leftPart)?"left":"right"},disableRestoreFocus:!0,children:Object(a.jsx)(f.a,{m:1,children:Object(a.jsx)(Ae,{data:M,time:null===C||void 0===C?void 0:C.date})})}),Object(a.jsx)("svg",{width:"100%",height:r,children:Object(a.jsxs)("g",{transform:"translate(".concat(p,", ").concat(O,")"),children:[Object(a.jsx)("defs",{children:Object(a.jsx)("clipPath",{id:"clip-line",children:Object(a.jsx)("rect",{transform:"translate(0, -2)",width:y.range()[1]+4,height:_.range()[0]+2})})}),Object(a.jsx)(ve,{xScale:y,height:m}),Object(a.jsx)(ge,{yScale:_,label:""}),t.map((function(e,t){return Object(a.jsx)("path",{stroke:i(e.metadata.name),className:"line",style:{opacity:void 0!==(null===C||void 0===C?void 0:C.activeSeries)?t===(null===C||void 0===C?void 0:C.activeSeries)?1:.2:1},d:U(e),clipPath:"url(#clip-line)"},t)})),Object(a.jsx)("g",{ref:I,children:Object(a.jsx)(Se,{height:m,x:null===C||void 0===C?void 0:C.xCoord})}),Object(a.jsx)(Ee,{xScale:y,yScale:_,datesInChart:P,onInteraction:D,setSelection:L})]})})]})}})},Re=n(113),_e=function(e){if(0===Object.keys(e.metric).length)return"Query result";var t=e.metric,n=t.__name__,a=Object(Re.a)(t,["__name__"]);return"".concat(n||""," {").concat(Object.entries(a).map((function(e){return"".concat(e[0],": ").concat(e[1])})).join(", "),"}")},ke=n(421),Ue=Object(X.a)({legendWrapper:{display:"grid",width:"100%",gridTemplateColumns:"repeat(auto-fit)",gridColumnGap:".5em",paddingLeft:"8px"}}),De=function(e){var t=e.labels,n=e.onChange,r=e.categories,i=Ue(),o=Object(c.useMemo)((function(){return t.length>0?r.filter((function(e){return 1===e.variations})).map((function(e){return"".concat(e.key,": ").concat(t[0].labelData[e.key])})):[]}),[r,t]),s=Object(c.useMemo)((function(){return r.filter((function(e){return 1!==e.variations})).map((function(e){return e.key}))}),[r]);return Object(a.jsxs)("div",{children:[Object(a.jsx)("div",{style:{textAlign:"center"},children:"Legend for ".concat(o.join(", "))}),Object(a.jsx)("div",{className:i.legendWrapper,children:t.map((function(e,t){return Object(a.jsx)("div",{children:Object(a.jsx)(y.a,{control:Object(a.jsx)(ke.a,{size:"small",checked:e.checked,onChange:function(){n(t)},style:{color:e.color,padding:"4px"}}),label:Object(a.jsx)(x.a,{variant:"body2",children:s.map((function(t){return"".concat(t,": ").concat(e.labelData[t])})).join(", ")})})},e.seriesName)}))})]})},Me=function(e){return Object(c.useMemo)((function(){var t={};return e.forEach((function(e){return Object.entries(e.metric).forEach((function(e){return t[e[0]]?t[e[0]].options.add(e[1]):t[e[0]]={options:new Set([e[1]])}}))})),Object.entries(t).map((function(e){return{key:e[0],variations:e[1].options.size}})).sort((function(e,t){return e.variations-t.variations}))}),[e])},Ie=Object(X.a)({inlineBtn:{"&:hover":{cursor:"pointer"}}}),Ne=function(e){var t=e.handler,n=e.text,c=Ie();return Object(a.jsx)(p.a,{component:"span",className:c.inlineBtn,onClick:t,children:n})},Pe=Oe.n,Le=function(e){var t=e.data,n=e.timePresets,r=Object(c.useState)(20),i=Object(s.a)(r,2),o=i[0],l=i[1],u=Object(c.useMemo)((function(){return null===t||void 0===t?void 0:t.map((function(e){return{metadata:{name:_e(e)},metric:e.metric,values:e.values.map((function(e){return{key:e[0],value:+e[1]}}))}}))}),[t]),j=Object(c.useMemo)((function(){return u.slice(0,o)}),[u,o]),d=Me(t),b=Object(c.useMemo)((function(){return j.map((function(e){return e.metadata.name}))}),[j]),O=Object(c.useState)(b),h=Object(s.a)(O,2),f=h[0],x=h[1];Object(c.useEffect)((function(){f.join(",")!==b.join(",")&&x(b)}),[b,x,f]);var p=Object(c.useMemo)((function(){return u.length}),[u]),v=Object(c.useMemo)((function(){var e=f.length,t=e<=Pe.length?Pe:Object(Oe.j)(e).map((function(t){return t/e})).map(Oe.e);return Object(Oe.l)().domain(f).range(t)}),[f]),m=Object(c.useMemo)((function(){return f.map((function(e){var t;return{color:v(e),seriesName:e,labelData:null===(t=j.find((function(t){return t.metadata.name===e})))||void 0===t?void 0:t.metric,checked:!0}}))}),[v,f]),y=Object(c.useState)(m),g=Object(s.a)(y,2),S=g[0],E=g[1];Object(c.useEffect)((function(){E(m)}),[m]);var T=Object(c.useMemo)((function(){return S.filter((function(e){return e.checked})).map((function(e){return e.seriesName}))}),[S]),w=Object(c.useMemo)((function(){return j.filter((function(e){return T.includes(e.metadata.name)}))}),[j,T]);return Object(a.jsx)(a.Fragment,{children:p>0?Object(a.jsxs)(a.Fragment,{children:[p>20&&Object(a.jsx)("div",{style:{textAlign:"center"},children:p>o?Object(a.jsxs)("span",{style:{fontStyle:"italic"},children:["Showing only first ",o," of ",p," series.\xa0",o+20>=p?Object(a.jsx)(Ne,{handler:function(){return l(p)},text:"Show all"}):Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)(Ne,{handler:function(){return l((function(e){return Math.min(e+20,p)}))},text:"Show ".concat(20," more")}),",\xa0",Object(a.jsx)(Ne,{handler:function(){return l(p)},text:"show all"}),"."]})]}):Object(a.jsxs)("span",{style:{fontStyle:"italic"},children:["Showing all series.\xa0",Object(a.jsx)(Ne,{handler:function(){return l(20)},text:"Show only ".concat(20)}),"."]})}),Object(a.jsx)(Ce,{height:400,series:w,color:v,timePresets:n,categories:d}),Object(a.jsx)(De,{labels:S,onChange:function(e){E((function(t){if(t){var n=Object(be.a)(t);return n[e]=Object(C.a)(Object(C.a)({},n[e]),{},{checked:!n[e].checked}),n}return t}))},categories:d})]}):Object(a.jsx)("div",{style:{textAlign:"center"},children:"No data to show"})})},He=n(391),Fe=n(204),Be=n(392),ze=n(393),Ye=n(394),qe=n(395),Qe=n(396),We=Object(X.a)({deemphasized:{opacity:.4}}),Ge=function(e){var t=e.data,n=We(),r=Me(t),i=Object(c.useMemo)((function(){return null===t||void 0===t?void 0:t.map((function(e){return{metadata:r.map((function(t){return e.metric[t.key]||"-"})),value:e.value[1]}}))}),[r,t]);return Object(a.jsx)(a.Fragment,{children:i.length>0?Object(a.jsx)(He.a,{component:Fe.a,children:Object(a.jsxs)(Be.a,{"aria-label":"simple table",children:[Object(a.jsx)(ze.a,{children:Object(a.jsxs)(Ye.a,{children:[r.map((function(e,t){return Object(a.jsx)(qe.a,{style:{textTransform:"capitalize"},children:e.key},t)})),Object(a.jsx)(qe.a,{align:"right",children:"Value"})]})}),Object(a.jsx)(Qe.a,{children:i.map((function(e,t){return Object(a.jsxs)(Ye.a,{children:[e.metadata.map((function(e,c){var r=i[t-1]&&i[t-1].metadata[c];return Object(a.jsx)(qe.a,{className:r===e?n.deemphasized:void 0,children:e},c)})),Object(a.jsx)(qe.a,{align:"right",children:e.value})]},t)}))})]})}):Object(a.jsx)("div",{style:{textAlign:"center"},children:"No data to show"})})},Je=n(423),Ve=n(410),Ze=n(411),$e=n(407),Ke=n(417),Xe=n(3),et=n(9),tt=n(109),nt=n(200),at=n(201),ct={windows:"Windows",mac:"Mac OS",linux:"Linux"},rt=function(){return(Object.values(ct).find((function(e){return navigator.userAgent.indexOf(e)>=0}))||"unknown")===ct.mac},it=function(e){var t=e.query,n=e.setQuery,r=e.runQuery,i=e.server,o=e.oneLiner,l=void 0!==o&&o,u=e.autocomplete,j=Object(c.useRef)(null),d=Object(c.useState)(),b=Object(s.a)(d,2),O=b[0],h=b[1];return Object(c.useEffect)((function(){return j.current&&h(new et.d({parent:j.current})),function(){return null===O||void 0===O?void 0:O.destroy()}}),[]),Object(c.useEffect)((function(){var e=new nt.a;e.activateCompletion(u),e.setComplete({url:i});var a=et.d.updateListener.of((function(e){e.docChanged&&n(e.state.doc.toJSON().map((function(e){return e.trim()})).join(""))}));null===O||void 0===O||O.setState(Xe.e.create({doc:t,extensions:[at.a,Object(et.k)(tt.a),a,e.asExtension(),Object(et.k)([{key:rt()?"Cmd-Enter":"Ctrl-Enter",run:function(){return r(),!0}}])]}))}),[i,O,u]),Object(a.jsx)(a.Fragment,{children:Object(a.jsx)("div",{ref:j,className:l?"one-line-scroll":void 0})})},ot=n(415),st=function(){return Object(a.jsx)(He.a,{component:Fe.a,children:Object(a.jsxs)(Be.a,{"aria-label":"simple table",size:"small",children:[Object(a.jsx)(ze.a,{children:Object(a.jsxs)(Ye.a,{children:[Object(a.jsx)(qe.a,{children:"Long"}),Object(a.jsx)(qe.a,{children:"Short"})]})}),Object(a.jsx)(Qe.a,{children:D.map((function(e,t){return Object(a.jsxs)(Ye.a,{children:[Object(a.jsx)(qe.a,{component:"th",scope:"row",children:e.long}),Object(a.jsx)(qe.a,{children:e.short})]},t)}))})]})})},lt=function(e){var t=e.setDuration,n=Object(c.useState)(!1),i=Object(s.a)(n,2),o=i[0],l=i[1],u=r.a.useState(null),j=Object(s.a)(u,2),d=j[0],b=j[1],O=Object(c.useState)(),h=Object(s.a)(O,2),p=h[0],v=h[1],m=J().time,y=m.period.end,g=m.duration,S=V(),E=Object(c.useState)(g),T=Object(s.a)(E,2),w=T[0],A=T[1];Object(c.useEffect)((function(){A(g)}),[g]),Object(c.useEffect)((function(){var e;v((e=N(y),_()(e).format("YYYY-MM-DD[T]HH:mm:ss")))}),[y]),Object(c.useEffect)((function(){o||t(w)}),[w,o]);var C=function(){b(null)},R=Boolean(d);return Object(a.jsxs)(f.a,{m:1,flexDirection:"row",display:"flex",children:[Object(a.jsxs)(f.a,{px:1,children:[Object(a.jsx)(f.a,{children:Object(a.jsx)(Ke.a,{label:"Duration",value:w,onChange:function(e){A(e.target.value)},fullWidth:!0,onBlur:function(){l(!1)},onFocus:function(){l(!0)}})}),Object(a.jsx)(f.a,{my:2,children:Object(a.jsxs)(x.a,{variant:"body2",children:["Possible options",Object(a.jsx)("span",{"aria-owns":R?"mouse-over-popover":void 0,"aria-haspopup":"true",style:{cursor:"pointer"},onMouseEnter:function(e){b(e.currentTarget)},onMouseLeave:C,children:":\xa0"}),Object(a.jsx)(Te.a,{open:R,anchorEl:d,anchorOrigin:{vertical:"bottom",horizontal:"left"},transformOrigin:{vertical:"top",horizontal:"left"},style:{pointerEvents:"none"},onClose:C,disableRestoreFocus:!0,children:Object(a.jsx)(st,{})}),Object(a.jsx)(Ne,{handler:function(){return A("5m")},text:"5m"}),",\xa0",Object(a.jsx)(Ne,{handler:function(){return A("1h")},text:"1h"}),",\xa0",Object(a.jsx)(Ne,{handler:function(){return A("1h 30m")},text:"1h 30m"})]})})]}),Object(a.jsxs)(f.a,{px:1,children:[Object(a.jsx)(f.a,{children:Object(a.jsx)(ot.a,{variant:"inline",ampm:!1,label:"Until",value:p,onChange:function(e){return S({type:"SET_UNTIL",payload:e})},onError:console.log,format:"DD/MM/YYYY HH:mm:ss"})}),Object(a.jsx)(f.a,{my:2,children:Object(a.jsxs)(x.a,{variant:"body2",children:["Will be changed to current time for auto-refresh mode.\xa0",Object(a.jsx)(Ne,{handler:function(){return S({type:"RUN_QUERY_TO_NOW"})},text:"Switch to now"})]})})]})]})},ut=n(191),jt=n.n(ut),dt=n(193),bt=n.n(dt),Ot=n(408),ht=n(428),ft=n(427),xt=n(405),pt=n(409),vt=n(419),mt=n(400),yt=n(398),gt=n(426),St=n(397),Et=n(403),Tt=n(399),wt=n(406),At=n(401),Ct=function(e){var t=e.children,n=e.value,c=e.index,r=Object(Re.a)(e,["children","value","index"]);return Object(a.jsx)("div",Object(C.a)(Object(C.a)({role:"tabpanel",hidden:n!==c,id:"auth-config-tabpanel-".concat(c),"aria-labelledby":"auth-config-tab-".concat(c)},r),{},{children:n===c&&Object(a.jsx)(f.a,{py:2,children:t})}))},Rt=n(189),_t=n.n(Rt),kt=n(190),Ut=n.n(kt),Dt={authMethod:"NO_AUTH",saveAuthLocally:!1},Mt=L("AUTH_TYPE"),It=L("BASIC_AUTH_DATA"),Nt=L("BEARER_AUTH_DATA"),Pt=Object(C.a)(Object(C.a)({},Dt),{},{authMethod:Mt||Dt.authMethod,basicData:It,bearerData:Nt,saveAuthLocally:!(!It&&!Nt)}),Lt=function(){H(F)};function Ht(e,t){switch(t.type){case"SET_BASIC_AUTH":return t.payload.checkbox?P("BASIC_AUTH_DATA",t.payload.value):Lt(),P("AUTH_TYPE","BASIC_AUTH"),Object(C.a)(Object(C.a)({},e),{},{authMethod:"BASIC_AUTH",basicData:t.payload.value});case"SET_BEARER_AUTH":return t.payload.checkbox?P("BEARER_AUTH_DATA",t.payload.value):Lt(),P("AUTH_TYPE","BEARER_AUTH"),Object(C.a)(Object(C.a)({},e),{},{authMethod:"BEARER_AUTH",bearerData:t.payload.value});case"SET_NO_AUTH":return!t.payload.checkbox&&Lt(),P("AUTH_TYPE","NO_AUTH"),Object(C.a)(Object(C.a)({},e),{},{authMethod:"NO_AUTH"});default:throw new Error}}var Ft=Object(c.createContext)({}),Bt=function(){return Object(c.useContext)(Ft).state},zt=function(e){var t=e.children,n=Object(c.useReducer)(Ht,Pt),r=Object(s.a)(n,2),i=r[0],o=r[1],l=Object(c.useMemo)((function(){return{state:i,dispatch:o}}),[i,o]);return Object(a.jsx)(Ft.Provider,{value:l,children:t})},Yt=Object(X.a)((function(){return Object(ft.a)({tabsContent:{height:"200px"}})})),qt="Bearer ",Qt=[{title:"No auth",id:"NO_AUTH"},{title:"Basic Auth",id:"BASIC_AUTH"},{title:"Bearer Token",id:"BEARER_AUTH"}],Wt=function(e){var t=Yt(),n=e.onClose,r=e.open,i=Bt(),o=i.saveAuthLocally,l=i.basicData,u=i.bearerData,j=i.authMethod,d=Object(c.useContext)(Ft).dispatch,b=Object(c.useState)(o),O=Object(s.a)(b,2),h=O[0],p=O[1],v=Object(c.useState)(l||{password:"",login:""}),m=Object(s.a)(v,2),g=m[0],S=m[1],E=Object(c.useState)((null===u||void 0===u?void 0:u.token)||qt),T=Object(s.a)(E,2),w=T[0],A=T[1],R=Object(c.useState)(Qt.findIndex((function(e){return e.id===j}))||0),_=Object(s.a)(R,2),k=_[0],U=_[1],D=function(){n()};return Object(a.jsxs)(ht.a,{onClose:D,"aria-labelledby":"simple-dialog-title",open:r,children:[Object(a.jsx)(Ot.a,{id:"simple-dialog-title",children:"Request Auth Settings"}),Object(a.jsxs)(xt.a,{children:[Object(a.jsx)(pt.a,{children:"This affects Authorization header sent to the server you specify. Not shown in URL and can be optionally stored on a client side"}),Object(a.jsx)(vt.a,{value:k,onChange:function(e,t){U(t)},indicatorColor:"primary",textColor:"primary",children:Qt.map((function(e){return Object(a.jsx)(mt.a,{label:e.title},e.id)}))}),Object(a.jsxs)(f.a,{p:0,display:"flex",flexDirection:"column",className:t.tabsContent,children:[Object(a.jsxs)(f.a,{flexGrow:1,children:[Object(a.jsx)(Ct,{value:k,index:0,children:Object(a.jsx)(x.a,{style:{fontStyle:"italic"},children:"No Authorization Header"})}),Object(a.jsxs)(Ct,{value:k,index:1,children:[Object(a.jsxs)(yt.a,{margin:"dense",fullWidth:!0,children:[Object(a.jsx)(gt.a,{htmlFor:"basic-login",children:"User"}),Object(a.jsx)(St.a,{id:"basic-login",startAdornment:Object(a.jsx)(Et.a,{position:"start",children:Object(a.jsx)(_t.a,{})}),required:!0,onChange:function(e){return S((function(t){return Object(C.a)(Object(C.a)({},t),{},{login:e.target.value||""})}))},value:(null===g||void 0===g?void 0:g.login)||""})]}),Object(a.jsxs)(yt.a,{margin:"dense",fullWidth:!0,children:[Object(a.jsx)(gt.a,{htmlFor:"basic-pass",children:"Password"}),Object(a.jsx)(St.a,{id:"basic-pass",startAdornment:Object(a.jsx)(Et.a,{position:"start",children:Object(a.jsx)(Ut.a,{})}),onChange:function(e){return S((function(t){return Object(C.a)(Object(C.a)({},t),{},{password:e.target.value||""})}))},value:(null===g||void 0===g?void 0:g.password)||""})]})]}),Object(a.jsx)(Ct,{value:k,index:2,children:Object(a.jsx)(Ke.a,{id:"bearer-auth",label:"Bearer token",multiline:!0,fullWidth:!0,value:w,onChange:function(e){var t=e.target.value;t.startsWith(qt)?A(t):A(qt)},InputProps:{onPaste:function(e){var t=e.clipboardData.getData("text/plain");t.startsWith(qt)?A(t):A(qt+t),e.preventDefault()}},rowsMax:6})})]}),Object(a.jsxs)(yt.a,{children:[Object(a.jsx)(y.a,{control:Object(a.jsx)(ke.a,{checked:h,onChange:function(){return p((function(e){return!e}))},name:"checkedB",color:"primary"}),label:"Persist Auth Data Locally"}),Object(a.jsx)(Tt.a,{children:h?"Auth Data and the Selected method will be saved to LocalStorage":"Auth Data won't be saved. All previously saved Auth Data will be removed"})]})]})]}),Object(a.jsx)(wt.a,{children:Object(a.jsx)(At.a,{onClick:function(){switch(k){case 0:d({type:"SET_NO_AUTH",payload:{checkbox:h}});break;case 1:d({type:"SET_BASIC_AUTH",payload:{checkbox:h,value:g}});break;case 2:d({type:"SET_BEARER_AUTH",payload:{checkbox:h,value:{token:w}}})}D()},color:"primary",children:"Apply"})})]})},Gt=n(192),Jt=n.n(Gt),Vt=n(387),Zt=n(194),$t=n.n(Zt),Kt=function(){var e=J(),t=e.serverUrl,n=e.query,r=e.time.duration,i=V(),o=J().queryControls.autocomplete,l=Object(c.useState)(!1),u=Object(s.a)(l,2),j=u[0],d=u[1],b=Object(c.useState)(!0),O=Object(s.a)(b,2),h=O[0],p=O[1],v=Object(c.useState)(!1),m=Object(s.a)(v,2),T=m[0],w=m[1],A=Object(c.useRef)(null),C=Object(c.useRef)(null),R=function(){return i({type:"RUN_QUERY"})};return Object(a.jsxs)(a.Fragment,{children:[Object(a.jsxs)(Je.a,{expanded:h,onChange:function(){return p((function(e){return!e}))},children:[Object(a.jsxs)(Ve.a,{expandIcon:Object(a.jsx)(jt.a,{}),"aria-controls":"panel1a-content",id:"panel1a-header",children:[Object(a.jsx)(f.a,{mr:2,children:Object(a.jsx)(x.a,{variant:"h6",component:"h2",children:"Query Configuration"})}),Object(a.jsx)(f.a,{flexGrow:1,onClick:function(e){return e.stopPropagation()},onFocusCapture:function(e){return e.stopPropagation()},children:Object(a.jsx)(Vt.a,{disablePortal:!h,container:C.current,children:Object(a.jsx)(it,{server:t,query:n,oneLiner:!h,autocomplete:o,runQuery:R,setQuery:function(e){return i({type:"SET_QUERY",payload:e})}})})})]}),Object(a.jsx)(Ze.a,{children:Object(a.jsxs)($e.a,{container:!0,spacing:2,children:[Object(a.jsx)($e.a,{item:!0,xs:12,md:6,children:Object(a.jsxs)(f.a,{children:[Object(a.jsxs)(f.a,{py:2,display:"flex",alignItems:"center",children:[Object(a.jsx)(Ke.a,{variant:"outlined",fullWidth:!0,label:"Server URL",value:t,inputProps:{style:{fontFamily:"Monospace"}},onChange:function(e){var t=e.target.value;i({type:"SET_SERVER",payload:t})}}),Object(a.jsx)(f.a,{ml:1,children:Object(a.jsx)(S.a,{title:"Execute Query",children:Object(a.jsx)(E.a,{onClick:R,children:Object(a.jsx)(Jt.a,{})})})}),Object(a.jsx)(f.a,{children:Object(a.jsx)(S.a,{title:"Request Auth Settings",children:Object(a.jsx)(E.a,{onClick:function(){return d(!0)},children:Object(a.jsx)(bt.a,{})})})})]}),Object(a.jsxs)(f.a,{py:2,display:"flex",children:[Object(a.jsx)(f.a,{flexGrow:1,mr:2,children:Object(a.jsx)("div",{ref:C})}),Object(a.jsxs)("div",{children:[Object(a.jsx)(S.a,{title:"Query Editor Settings",children:Object(a.jsx)(E.a,{onClick:function(){return w(!T)},children:Object(a.jsx)($t.a,{ref:A})})}),Object(a.jsx)(Te.a,{open:T,transformOrigin:{vertical:-20,horizontal:"left"},onClose:function(){return w(!1)},anchorEl:A.current,children:Object(a.jsx)(f.a,{p:2,children:Object(a.jsx)(y.a,{control:Object(a.jsx)(g.a,{size:"small",checked:o,onChange:function(){i({type:"TOGGLE_AUTOCOMPLETE"}),P("AUTOCOMPLETE",!o)}}),label:"Autocomplete"})})})]})]})]})}),Object(a.jsx)($e.a,{item:!0,xs:8,md:6,children:Object(a.jsx)(f.a,{style:{borderRadius:"4px",borderColor:"#b9b9b9",borderStyle:"solid",borderWidth:"1px",height:"calc(100% - 18px)",marginTop:"16px"},children:Object(a.jsx)(lt,{setDuration:function(e){return i({type:"SET_DURATION",payload:e})},duration:r})})})]})})]}),Object(a.jsx)(Wt,{open:j,onClose:function(){return d(!1)}})]})},Xt=function(){var e=J(),t=e.query,n=e.displayType,a=e.serverUrl,r=e.time.period,i=Bt(),o=i.basicData,l=i.bearerData,u=i.authMethod,j=Object(c.useState)(!1),d=Object(s.a)(j,2),b=d[0],O=d[1],h=Object(c.useState)(),f=Object(s.a)(h,2),x=f[0],p=f[1],v=Object(c.useState)(),m=Object(s.a)(v,2),y=m[0],g=m[1],S=Object(c.useState)(),E=Object(s.a)(S,2),T=E[0],w=E[1];Object(c.useEffect)((function(){T&&(p(void 0),g(void 0))}),[T]);var A=Object(c.useMemo)((function(){if(r){if(!a)return void w("Please enter Server URL");if(!t.trim())return void w("Please enter a valid Query and execute it");if(function(e){var t;try{t=new URL(e)}catch(n){return!1}return"http:"===t.protocol||"https:"===t.protocol}(a))return"chart"===n?function(e,t,n){return"".concat(e,"/api/v1/query_range?query=").concat(encodeURIComponent(t),"&start=").concat(n.start,"&end=").concat(n.end,"&step=").concat(n.step)}(a,t,r):function(e,t,n){return"".concat(e,"/api/v1/query?query=").concat(encodeURIComponent(t),"&start=").concat(n.start,"&end=").concat(n.end,"&step=").concat(n.step)}(a,t,r);w("Please provide a valid URL")}}),[a,r,n]);return Object(c.useEffect)((function(){Object(xe.a)(fe.a.mark((function e(){var c,r,i,s;return fe.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!A){e.next=33;break}return c=new Headers,"BASIC_AUTH"===u&&c.set("Authorization","Basic "+btoa("".concat((null===o||void 0===o?void 0:o.login)||"",":").concat((null===o||void 0===o?void 0:o.password)||""))),"BEARER_AUTH"===u&&c.set("Authorization",(null===l||void 0===l?void 0:l.token)||""),O(!0),e.next=7,fetch(A,{headers:c});case 7:if(!(r=e.sent).ok){e.next=18;break}return P("PREFERRED_URL",a),P("LAST_QUERY",t),e.next=13,r.json();case 13:i=e.sent,w(void 0),"chart"===n?p(i.data.result):g(i.data.result),e.next=32;break;case 18:return e.t0=w,e.next=21,r.json();case 21:if(e.t2=s=e.sent,e.t1=null===e.t2,e.t1){e.next=25;break}e.t1=void 0===s;case 25:if(!e.t1){e.next=29;break}e.t3=void 0,e.next=30;break;case 29:e.t3=s.error;case 30:e.t4=e.t3,(0,e.t0)(e.t4);case 32:O(!1);case 33:case"end":return e.stop()}}),e)})))()}),[A,a,n]),{fetchUrl:A,isLoading:b,graphData:x,liveData:y,error:T}},en=function(e){var t=e.data,n=d().showInfoMessage,r=Object(c.useMemo)((function(){return JSON.stringify(t,null,2)}),[t]);return Object(a.jsxs)(f.a,{position:"relative",children:[Object(a.jsx)(f.a,{flexDirection:"column",justifyContent:"flex-end",display:"flex",style:{position:"fixed",right:"16px"},children:Object(a.jsx)(At.a,{variant:"outlined",onClick:function(e){navigator.clipboard.writeText(r),n("Formatted JSON has been copied"),e.preventDefault()},children:"Copy JSON"})}),Object(a.jsx)("pre",{children:r})]})},tn=n(195),nn=n.n(tn),an=function(e){var t=e.url,n=d().showInfoMessage;return Object(a.jsx)(f.a,{pl:2,py:1,flexShrink:0,display:"flex",children:Object(a.jsx)(S.a,{title:"Copy Query URL",children:Object(a.jsx)(E.a,{size:"small",onClick:function(e){t&&(navigator.clipboard.writeText(t),n("Value has been copied"),e.preventDefault())},children:Object(a.jsx)(nn.a,{style:{color:"white"}})})})})},cn=function(){var e=J(),t=e.displayType,n=e.time.period,c=Xt(),r=c.fetchUrl,i=c.isLoading,o=c.liveData,s=c.graphData,l=c.error;return Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)(O.a,{position:"static",children:Object(a.jsxs)(h.a,{children:[Object(a.jsxs)(f.a,{display:"flex",children:[Object(a.jsxs)(x.a,{variant:"h5",children:[Object(a.jsx)("span",{style:{fontWeight:"bolder"},children:"VM"}),Object(a.jsx)("span",{style:{fontWeight:"lighter"},children:"UI"})]}),Object(a.jsx)("div",{style:{fontSize:"10px",marginTop:"-2px"},children:Object(a.jsx)("div",{children:"BETA"})})]}),Object(a.jsx)("div",{style:{fontSize:"10px",position:"absolute",top:"40px",opacity:".4"},children:Object(a.jsx)(p.a,{color:"inherit",href:"https://github.com/VictoriaMetrics/vmui/issues/new",target:"_blank",children:"Create an issue"})}),Object(a.jsx)(f.a,{ml:4,flexGrow:1,children:Object(a.jsx)(te,{})}),Object(a.jsx)(de,{}),Object(a.jsx)(an,{url:r})]})}),Object(a.jsxs)(f.a,{display:"flex",flexDirection:"column",style:{minHeight:"calc(100vh - 64px)"},children:[Object(a.jsx)(f.a,{m:2,children:Object(a.jsx)(Kt,{})}),Object(a.jsxs)(f.a,{flexShrink:1,children:[i&&Object(a.jsx)(v.a,{in:i,style:{transitionDelay:i?"300ms":"0ms"},children:Object(a.jsx)(f.a,{alignItems:"center",flexDirection:"column",display:"flex",style:{width:"100%",maxWidth:"calc(100vh - 32px)",position:"absolute",height:"150px",background:"linear-gradient(rgba(255,255,255,.7), rgba(255,255,255,.7), rgba(255,255,255,0))"},m:2,children:Object(a.jsx)(m.a,{})})}),Object(a.jsxs)(f.a,{p:2,children:[l&&Object(a.jsx)(u.a,{color:"error",style:{fontSize:"14px"},children:l}),s&&n&&"chart"===t&&Object(a.jsx)(Le,{data:s,timePresets:n}),o&&"code"===t&&Object(a.jsx)(en,{data:o}),o&&"table"===t&&Object(a.jsx)(Ge,{data:o})]})]})]})]})},rn=n(199),on=n(414),sn=n(413),ln=n(31),un=n(196),jn=function(){var e=Object(rn.a)({typography:{fontSize:10}});return Object(a.jsxs)(a.Fragment,{children:[Object(a.jsx)(sn.a,{})," ",Object(a.jsxs)(ln.a,{utils:un.a,children:[" ",Object(a.jsxs)(on.a,{theme:e,children:["  ",Object(a.jsxs)($,{children:[" ",Object(a.jsxs)(zt,{children:[" ",Object(a.jsxs)(b,{children:[" ",Object(a.jsx)(cn,{})]})]})]})]})]})]})},dn=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,431)).then((function(t){var n=t.getCLS,a=t.getFID,c=t.getFCP,r=t.getLCP,i=t.getTTFB;n(e),a(e),c(e),r(e),i(e)}))};o.a.render(Object(a.jsx)(r.a.StrictMode,{children:Object(a.jsx)(jn,{})}),document.getElementById("root")),dn()}},[[349,1,2]]]);