import request from '/@/utils/request'
// 查询科室列表
export function listDepartment(query:object) {
  return request({
    url: '/api/v1/archive/department/list',
    method: 'get',
    params: query
  })
}
// 查询科室详细
export function getDepartment(deptId:number) {
  return request({
    url: '/api/v1/archive/department/get',
    method: 'get',
    params: {
      deptId: deptId.toString()
    }
  })
}
// 新增科室
export function addDepartment(data:object) {
  return request({
    url: '/api/v1/archive/department/add',
    method: 'post',
    data: data
  })
}
// 修改科室
export function updateDepartment(data:object) {
  return request({
    url: '/api/v1/archive/department/edit',
    method: 'put',
    data: data
  })
}
// 删除科室
export function delDepartment(deptIds:number[]) {
  return request({
    url: '/api/v1/archive/department/delete',
    method: 'delete',
    data:{
      deptIds:deptIds
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
