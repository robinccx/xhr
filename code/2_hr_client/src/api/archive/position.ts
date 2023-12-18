import request from '/@/utils/request'
// 查询职务列表
export function listPosition(query:object) {
  return request({
    url: '/api/v1/archive/position/list',
    method: 'get',
    params: query
  })
}
// 查询职务详细
export function getPosition(positionId:number) {
  return request({
    url: '/api/v1/archive/position/get',
    method: 'get',
    params: {
      positionId: positionId.toString()
    }
  })
}
// 新增职务
export function addPosition(data:object) {
  return request({
    url: '/api/v1/archive/position/add',
    method: 'post',
    data: data
  })
}
// 修改职务
export function updatePosition(data:object) {
  return request({
    url: '/api/v1/archive/position/edit',
    method: 'put',
    data: data
  })
}
// 删除职务
export function delPosition(positionIds:number[]) {
  return request({
    url: '/api/v1/archive/position/delete',
    method: 'delete',
    data:{
      positionIds:positionIds
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
