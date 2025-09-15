const time  = document.getElementById('time');
const date  = document.getElementById('date');
const greet = document.getElementById('greet');

const quotes = [
  '抬头见月，侧耳听枫',
  '与你无关的事，别问，别想，别多嘴。 这个世界是残酷的但也美好的',
  '愿你被世界温柔以待，也能温柔地对待这个世界',
  '我们仰望同一片天空，却看着不同的地方',
  '生存的意志是最强的！',
  '虚伪的眼泪会伤害别人，虚伪的笑容会伤害自己',
  '大部分人都不想长大，只是没办法继续当一个小孩子'

];
function pickQuote() {
  const idx = Math.floor(Math.random() * quotes.length);
  return quotes[idx];}



function tick(){
    const now=new Date();
     time.textContent = now.toLocaleTimeString('zh-CN');

  date.textContent = now.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });

  if (!greet.dataset.done) {
    greet.textContent = pickQuote();
    greet.dataset.done = 'true';
  }
}
tick();
setInterval(tick,1000);

const $ = id => document.getElementById(id);
const slogans = ['南大家园','云家园','家园工作室','小家园传声机'];
const randBtn = $('randBtn'); 
randBtn.onclick = () => {
  const randomText = slogans[Math.floor(Math.random() * slogans.length)];
  $('search').value = randomText;
};

const engineMap = {
  google: 'https://www.google.com/search?q=',
  bing:  'https://cn.bing.com/search?q=',
};

const input  = document.getElementById('search'); 
const btn    = document.getElementById('searchBtn');
const select = document.getElementById('engine');

function doSearch() {
  const raw =input.value;
  const key = raw.trim();

  if (!key || raw === '有问题就快搜索吧...') {
    input.value = '请输入内容';
    input.select();
    return;
  }
  const base = engineMap[select.value];

  window.open(base + encodeURIComponent(key), '_blank');
  localStorage.setItem('lastSearch', key)

}


const search = document.getElementById('search');
document.getElementById('searchBtn').addEventListener('click', function () {
  const keyword = search.value.trim();
if (keyword) {
    localStorage.setItem('lastSearch', keyword)
  }
});

btn.addEventListener('click', doSearch);
input.addEventListener('keydown', e => {
  if (e.key === 'Enter') doSearch();
});


document.getElementById('prevBtn').addEventListener('click', function () {
  const last = localStorage.getItem('lastSearch');
  if (last) {
    search.value = last;

  }
});




