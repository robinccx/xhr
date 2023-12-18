import request from '/@/utils/request'
// 月度入职统计
export function ListMonthlyTotal(query:object) {
  return request({
    url: '/api/v1/report/nurse/monthlytotal',
    method: 'get',
    params: query
  })
}

// 月度职称统计
export function ListMonthlyTitle(query:object) {
    return request({
      url: '/api/v1/report/nurse/monthlytitle',
      method: 'get',
      params: query
    })
  }

  // 月度学历统计
export function ListMonthlyEdu(query:object) {
    return request({
      url: '/api/v1/report/nurse/monthlyedu',
      method: 'get',
      params: query
    })
  }

  // 季度填报数据统计
export function ListQuarterTotal(query:object) {
    return request({
      url: '/api/v1/report/nurse/quartertotal',
      method: 'get',
      params: query
    })
  }

  
  // 获取首页数据
export function ListHome(query:object) {
  return request({
    url: '/api/v1/report/nurse/home',
    method: 'get',
    params: query
  })
}