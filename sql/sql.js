module.exports = {
    insert: 'INSERT INTO movielist SET ?',
    queryAll: 'SELECT * FROM movielist',
    queryByKind: 'SELECT * FROM movielist WHERE kind = ? limit ?, ?',
    updateByName: 'UPDATE movielist SET ? WHERE name = ?',
    queryHotEmoji: 'SELECT * FROM hotemoji',
    queryByName: 'SELECT * FROM movielist WHERE name LIKE ',
    queryLogin: 'SELECT * FROM users WHERE UserName=?',
    queryArticle: 'SELECT * FROM article order by id desc limit ?,?',
    queryId: 'SELECT count(id) as count FROM users',
    queryCase: 'SELECT * FROM cases order by id desc limit ?,?',
    delUsers: 'DELETE FROM users WHERE Id=?',
    addUsers: 'INSERT INTO users (Id,ComTelephone,Password,UserName,CompanyId,IsAdmin) values (0,?,?,?,0,0)',
    putUsers: 'UPDATE users SET UserName=?,ComTelephone=?  WHERE Id=?',
    queryPass: 'SELECT Password FROM users WHERE UserName=?',
    putPass: 'UPDATE users SET Password=?  WHERE UserName=?',
    queryUsers: 'SELECT * FROM users order by id limit ?,?',
    delArticle: 'DELETE FROM article WHERE id=?',
    queryEssage: 'INSERT INTO article (id,title,tag,author,content,CreateTime,ArticleImg,see,chart,des) values (0,?,?,?,?,?,?,0,0,?)',
    artId: 'SELECT * FROM article WHERE id=?',
    caseId : 'SELECT * FROM cases WHERE id=?',
  novelKey: 'SELECT * FROM booktitles WHERE name LIKE ',// like后面的空格
  addNovel: 'INSERT INTO booktitles (id, author, name, titles) value (0, ?, ?, ?)'
};
