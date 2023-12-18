import request from '/@/utils/request'
// 查询职称列表
export function listTitle(query:object) {
  return request({
    url: '/api/v1/archive/title/list',
    method: 'get',
    params: query
  })
}
// 查询职称详细
export function getTitle(titleId:number) {
  return request({
    url: '/api/v1/archive/title/get',
    method: 'get',
    params: {
      titleId: titleId.toString()
    }
  })
}
// 新增职称
export function addTitle(data:object) {
  return request({
    url: '/api/v1/archive/title/add',
    method: 'post',
    data: data
  })
}
// 修改职称
export function updateTitle(data:object) {
  return request({
    url: '/api/v1/archive/title/edit',
    method: 'put',
    data: data
  })
}
// 删除职称
export function delTitle(titleIds:number[]) {
  return request({
    url: '/api/v1/archive/title/delete',
    method: 'delete',
    data:{
      titleIds:titleIds
    }
  })
}
//获取用户信息列表
export function getUserList(uIds:number[]){
    return request({
     url: '/api/v1/system/user/getUsers',
     method: 'get',
     params: {ids:uIds}
   })
}
