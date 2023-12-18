export interface DepartmentTableColumns {    
    deptId:number    
    deptCode:string;  // 科室代码    
    deptName:string;  // 科室名称    
    deptType:string;  // 片区    
    leader:string;  // 负责人    
    phone:string;  // 电话    
    email:string;  // 邮件    
    att1:string;  // 属性1    
    att2:string;  // 属性2    
    att3:string;  // 属性3    
    status:number;  // 状态    
    createdAt:string;  // 创建时间    
    createdBy:string;  // 创建用户    
}


export interface DepartmentInfoData {    
    deptId:number|undefined;        // 科室ID    
    parentId:number|undefined; // 父级ID    
    deptCode:string|undefined; // 科室代码    
    deptName:string|undefined; // 科室名称    
    deptType:string|undefined; // 片区    
    leader:string|undefined; // 负责人    
    phone:string|undefined; // 电话    
    email:string|undefined; // 邮件    
    att1:string|undefined; // 属性1    
    att2:string|undefined; // 属性2    
    att3:string|undefined; // 属性3    
    status:boolean; // 状态    
    createdAt:string|undefined; // 创建时间    
    createdBy:string|undefined; // 创建用户    
    updatedAt:string|undefined; //    
    updatedBy:string|undefined; // 修改用户    
    deletedAt:string|undefined; //    
    deletedBy:string|undefined; // 删除用户    
}


export interface DepartmentTableDataState {
    deptIds:any[];
    tableData: {
        data: Array<DepartmentTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            deptCode: string|undefined;            
            deptName: string|undefined;            
            status: number|undefined;            
            deptType: string|undefined;            
            att1: string|undefined;            
            att2: string|undefined;            
            att3: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface DepartmentEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:DepartmentInfoData;
    rules: object;
}