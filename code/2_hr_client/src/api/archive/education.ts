import request from '/@/utils/request'
// 查询学历列表
export function listEducation(query:object) {
  return request({
    url: '/api/v1/archive/education/list',
    method: 'get',
    params: query
  })
}
// 查询学历详细
export function getEducation(educationId:number) {
  return request({
    url: '/api/v1/archive/education/get',
    method: 'get',
    params: {
      educationId: educationId.toString()
    }
  })
}
// 新增学历
export function addEducation(data:object) {
  return request({
    url: '/api/v1/archive/education/add',
    method: 'post',
    data: data
  })
}
// 修改学历
export function updateEducation(data:object) {
  return request({
    url: '/api/v1/archive/education/edit',
    method: 'put',
    data: data
  })
}
// 删除学历
export function delEducation(educationIds:number[]) {
  return request({
    url: '/api/v1/archive/education/delete',
    method: 'delete',
    data:{
      educationIds:educationIds
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
