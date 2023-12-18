import request from '/@/utils/request'
// 查询代码生成测试列表
export function listDemoGen(query:object) {
  return request({
    url: '/api/v1/demo/demoGen/list',
    method: 'get',
    params: query
  })
}
// 查询代码生成测试详细
export function getDemoGen(id:number) {
  return request({
    url: '/api/v1/demo/demoGen/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增代码生成测试
export function addDemoGen(data:object) {
  return request({
    url: '/api/v1/demo/demoGen/add',
    method: 'post',
    data: data
  })
}
// 修改代码生成测试
export function updateDemoGen(data:object) {
  return request({
    url: '/api/v1/demo/demoGen/edit',
    method: 'put',
    data: data
  })
}
// 删除代码生成测试
export function delDemoGen(ids:number[]) {
  return request({
    url: '/api/v1/demo/demoGen/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 关联demo_gen_class表选项
export function listDemoGenClass(query:object){
   return request({
     url: '/api/v1/demo/demoGenClass/list',
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
