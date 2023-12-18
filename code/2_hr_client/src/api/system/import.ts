import request from '/@/utils/request'
// 查询导入列表
export function listNurseImport(query:object) {
  return request({
    url: '/api/v1/system/import/list',
    method: 'get',
    params: query
  })
}
// 查询导入详细
export function getNurseImport(nurseId:number) {
  return request({
    url: '/api/v1/system/import/get',
    method: 'get',
    params: {
      nurseId: nurseId.toString()
    }
  })
}
// 新增导入
export function addNurseImport(data:object) {
  return request({
    url: '/api/v1/system/import/add',
    method: 'post',
    data: data
  })
}
// 修改导入
export function updateNurseImport(data:object) {
  return request({
    url: '/api/v1/system/import/edit',
    method: 'put',
    data: data
  })
}
// 删除导入
export function delNurseImport(nurseIds:number[]) {
  return request({
    url: '/api/v1/system/import/delete',
    method: 'delete',
    data:{
      nurseIds:nurseIds
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
