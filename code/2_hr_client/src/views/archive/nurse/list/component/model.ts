export interface NurseTableColumns {    
    nurseId:number    
    nurseCode:string;  // 编号    
    nurseName:string;  // 名称    
    sex:number;  // 性别    
    birthday:string;  // 出生年月    
    areaId:number;  // 院区    
    deptId:number;  // 科室    
    startDate:string;  // 入职日期    
    educationId:number;  // 学历    
    titleId:number;  // 职称    
    staffType:number;  // 人员类别    
    note:string;  // 备注    
    endDate:string;  // 离职日期    
    status:number;  // 状态    
    sessionId:string;  //    
    certEnddate:string;  // 职业证书到期日期    
    regDate:string;  // 注册或变更日期    
    certNo:string;  // 职业证书    
    idNo:string;  // 身份证号    
    positionId:number;  // 职务    
    linkedNursePosition:LinkedNursePosition;    
    linkedNurseDepartment:LinkedNurseDepartment;    
    linkedNurseEducation:LinkedNurseEducation;    
    linkedNurseTitle:LinkedNurseTitle;    
}


export interface NurseInfoData {    
    nurseId:number|undefined;        // ID    
    nurseCode:string|undefined; // 编号    
    nurseName:string|undefined; // 名称    
    sex:number|undefined; // 性别    
    birthday:string|undefined; // 出生年月    
    areaId:number|undefined; // 院区    
    deptId:number|undefined; // 科室    
    linkedDeptId:LinkedNurseDepartment; // 科室    
    startDate:string|undefined; // 入职日期    
    educationId:number|undefined; // 学历    
    linkedEducationId:LinkedNurseEducation; // 学历    
    titleId:number|undefined; // 职称    
    linkedTitleId:LinkedNurseTitle; // 职称    
    staffType:number|undefined; // 人员类别    
    note:string|undefined; // 备注    
    endDate:string|undefined; // 离职日期    
    status:boolean; // 状态    
    createdAt:string|undefined; // 创建时间    
    createdBy:string|undefined; // 创建用户    
    updatedAt:string|undefined; //    
    updatedBy:string|undefined; // 修改用户    
    deletedAt:string|undefined; //    
    deletedBy:string|undefined; // 删除用户    
    sessionId:string|undefined; //    
    certEnddate:string|undefined; // 职业证书到期日期    
    regDate:string|undefined; // 注册或变更日期    
    certNo:string|undefined; // 职业证书    
    idNo:string|undefined; // 身份证号    
    positionId:number|undefined; // 职务    
    linkedPositionId:LinkedNursePosition; // 职务    
    linkedNursePosition:LinkedNursePosition;    
    linkedNurseDepartment:LinkedNurseDepartment;    
    linkedNurseEducation:LinkedNurseEducation;    
    linkedNurseTitle:LinkedNurseTitle;    
}


export interface LinkedNursePosition {	
    positionId:number|undefined;    //	
    positionName:string|undefined;    //	
}


export interface LinkedNurseDepartment {	
    deptId:number|undefined;    //	
    deptName:string|undefined;    //	
}


export interface LinkedNurseEducation {	
    educationId:number|undefined;    //	
    educationName:string|undefined;    //	
}


export interface LinkedNurseTitle {	
    titleId:number|undefined;    //	
    titleName:string|undefined;    //	
}


export interface NurseTableDataState {
    nurseIds:any[];
    tableData: {
        data: Array<NurseTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            areaId: number|undefined;            
            nurseName: string|undefined;            
            deptId: number|undefined;            
            status: number|undefined;            
            startDate: string[];            
            educationId: number|undefined;            
            titleId: number|undefined;            
            staffType: number|undefined;            
            nurseCode: string|undefined;            
            sessionId: string|undefined;            
            certEnddate: string|undefined;            
            regDate: string|undefined;            
            certNo: string|undefined;            
            idNo: string|undefined;            
            positionId: number|undefined;            
            dateRange: string[];
        };
    };
}


export interface NurseEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:NurseInfoData;
    rules: object;
}