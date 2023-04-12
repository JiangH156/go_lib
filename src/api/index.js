// 对所有的api接口进行统一管理
import requests from './request'

// 注册接口
export const register = (readerInfo) => requests({
    url: '/register',
    method: 'post',
    data: readerInfo
})
// 登录接口
export const login = (readerInfo) => requests({
    url: '/login',
    method: 'post',
    data: readerInfo
})
// 书籍接口
export const initBooksList = () => requests({
    url: '/books',
    method: 'post'
})
// 评论区接口
export const initCommentsList = () => requests({
    url: '/comments',
    method: 'post'
})
// 添加评论接口
export const addComment = (dataObj)=>requests({
	url: '/addcomment',
	method: 'post',
	data: dataObj
})
// 点赞接口
export const addPraise = (dataObj)=>requests({
	url: '/addpraise',
	method: 'post',
	data: dataObj
})
// 书名查找接口
export const searchBook = (bookNameObj) => requests({
    url: '/searchbook',
    method: 'post',
    data: bookNameObj
})



// 管理员查询借阅接口
export const initBorrowslist = () => requests({
    url: '/borrowslist',
    method: 'post'
})
export const initReportList = () => requests({
    url: '/initreportlist',
    method: 'post'
})
// 管理员查询预订接口
export const initReservelist = () => requests({
    url: '/reservelist',
    method: 'post'
})
// 管理员删除借阅记录接口
export const deleteBorrow = (borrowObj) => requests({
    url: '/deleteborrow',
    method: 'post',
    data:borrowObj
})
// 管理员通过读者ID查找借阅记录接口
export const searchBorrow = (infoObj) => requests({
    url: '/searchborrow',
    method: 'post',
    data:infoObj
})
// 管理员获取读者信息接口
export const initReaderList = () => requests({
    url: '/initreaderlist',
    method: 'post',
})
// 管理员添加书籍接口
export const addBooks = (infoObj) => requests({
    url: '/adminaddbooks',
    method: 'post',
    data:infoObj
})
// 管理员审核举报接口
export const auditComment = (infoObj) => requests({
    url: '/auditcomment',
    method: 'post',
    data:infoObj
})
// 管理员修改图书信息接口
export const changeBookInfo = (infoObj) => requests({
    url: '/changebookinfo',
    method: 'post',
    data:infoObj
})
// 管理员删除图书信息接口
export const delBook = (infoObj) => requests({
    url: '/delbook',
    method: 'post',
    data:infoObj
})
// 管理员删除人员信息接口
export const delPerson = (infoObj) => requests({
    url: '/delperson',
    method: 'post',
    data:infoObj
})
// 管理员提醒读者还书接口
export const alertPerson = (infoObj) => requests({
    url: '/alertperson',
    method: 'post',
    data:infoObj
})




// 读者请求借阅记录接口
export const initBorrows = (readerId) => requests({
    url: '/borrows',
    method: 'post',
    data: readerId
})
// 读者查询举报记录接口
export const initStuReport = (readerId) => requests({
    url: '/initstureport',
    method: 'post',
    data: readerId
})
// 添加预约记录接口
export const addReserve = (reserveObj) => requests({
    url: '/addreserve',
    method: 'post',
    data: reserveObj
})
// 删除预约记录接口
export const deleteReserve = (reserveObj) => requests({
    url: '/cancelreserve',
    method: 'post',
    data: reserveObj
})
// 借书接口
export const addBorrow = (borrowObj) => requests({
    url: '/addborrow',
    method: 'post',
    data: borrowObj
})
// 续借接口
export const continueBorrow = (infoObj) => requests({
    url: '/continueborrow',
    method: 'post',
    data: infoObj
})
// 还书接口
export const returnBook = (infoObj) => requests({
    url: '/returnbook',
    method: 'post',
    data: infoObj
})
// 查询预约接口
export const initReserve = (readerObj) => requests({
    url: '/reserve',
    method: 'post',
    data: readerObj
})
// 重新获取学生信息接口
export const initReader = (readerId) => requests({
    url: '/initreader',
    method: 'post',
    data: readerId
})
// 学生举报接口
export const reportComment = (infoObj)=>requests({
    url:'/reportcomment',
    method:'post',
    data:infoObj
})

