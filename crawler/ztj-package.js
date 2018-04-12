let cheerio = require('cheerio');
let request = require('request');
//let iconv = require('iconv-lite');
const express = require('express')
const app = express()
const superagent = require('superagent')
require('superagent-charset')(superagent)
const async = require('async');
function getHtmlByUrl(href,callback) {
    request(href.replace(/([\u4e00-\u9fa5])/g, (str) => encodeURIComponent(str) ), function(err, response, body) {
        if (!err && response.statusCode == 200) {
            //body即为目标页面的html
            let $ = cheerio.load(body);
            let v = $('.result-game-item-title-link').eq(0).attr('href');
            let v1 = $('.result-game-item-title-link').eq(10).attr('href');//undefined
            console.log(v,v1);
            superagent.get(v)
                .charset('gbk')
                .end(function (err, res) {
                    const $ = cheerio.load(res.text,{decodeEntities: false});
                    let content = [];
                    $('#list dd').each((i,v) => {
                        let $v = $(v)
                        content.push($v.find('a').text() + '+' + 'https://www.zwdu.com'+ $v.find('a').attr('href'))
                    })
                    let obj = {
                        name: $('#info h1').text(),
                        content: content.join('-'),
                        //urls: urls.join('-')
                    }
                    callback(obj)
                })
        } else {
            console.log('get page error url => ' + href);
        }
    });
}
//let url = 'https://www.zwdu.com/search.php?keyword=强人';
//URL编码，只转换中文
//getHtmlByUrl( url);
module.exports = getHtmlByUrl;
//可参考同目录下其他文件