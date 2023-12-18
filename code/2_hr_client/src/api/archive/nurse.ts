import request from '/@/utils/request'
// 查询护士列表
export function listNurse(query:object) {
  return request({
    url: '/api/v1/archive/nurse/list',
    method: 'get',
    params: query
  })
}
// 查询护士详细
export function getNurse(nurseId:number) {
  return request({
    url: '/api/v1/archive/nurse/get',
    method: 'get',
    params: {
      nurseId: nurseId.toString()
    }
  })
}
// 新增护士
export function addNurse(data:object) {
  return request({
    url: '/api/v1/archive/nurse/add',
    method: 'post',
    data: data
  })
}
// 修改护士
export function updateNurse(data:object) {
  return request({
    url: '/api/v1/archive/nurse/edit',
    method: 'put',
    data: data
  })
}
// 删除护士
export function delNurse(nurseIds:number[]) {
  return request({
    url: '/api/v1/archive/nurse/delete',
    method: 'delete',
    data:{
      nurseIds:nurseIds
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
// 关联t_position表选项
export function listPosition(query:object){
  return request({
    url: '/api/v1/archive/position/list',
    method: 'get',
    params: query
  })
}

// 关联t_education表选项
export function listEducation(query:object){
   return request({
     url: '/api/v1/archive/education/list',
     method: 'get',
     params: query
   })
}
// 关联t_title表选项
export function listTitle(query:object){
   return request({
     url: '/api/v1/archive/title/list',
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
// 导入护士
export function importNurse(data:object) {
  return request({
    url: '/api/v1/archive/nurse/import',
    method: 'post',
    data: data
  })
}

// 护士注册信息导入
export function updateByImport(data:object) {
  return request({
    url: '/api/v1/archive/nurse/updatebyimport',
    method: 'post',
    data: data
  })
}