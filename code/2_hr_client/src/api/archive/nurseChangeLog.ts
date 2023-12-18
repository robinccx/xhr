import request from '/@/utils/request'
// 查询变更记录查询列表
export function listNurseChangeLog(query:object) {
  return request({
    url: '/api/v1/archive/nurseChangeLog/list',
    method: 'get',
    params: query
  })
}
// 查询变更记录查询详细
export function getNurseChangeLog(id:number) {
  return request({
    url: '/api/v1/archive/nurseChangeLog/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增变更记录查询
export function addNurseChangeLog(data:object) {
  return request({
    url: '/api/v1/archive/nurseChangeLog/add',
    method: 'post',
    data: data
  })
}
// 修改变更记录查询
export function updateNurseChangeLog(data:object) {
  return request({
    url: '/api/v1/archive/nurseChangeLog/edit',
    method: 'put',
    data: data
  })
}
// 删除变更记录查询
export function delNurseChangeLog(ids:number[]) {
  return request({
    url: '/api/v1/archive/nurseChangeLog/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 关联t_department表选项
export function listDepartment(query:object){
   return request({
     url: '/api/v1/archive/department/list',
     method: 'get',
     params: query
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
