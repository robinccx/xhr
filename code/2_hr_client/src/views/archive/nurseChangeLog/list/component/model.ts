export interface NurseChangeLogTableColumns {    
    id:number    
    nurseId:number;  //    
    changeType:number;  // 变更类型    
    fromEntityName:string;  // 原内容    
    toEntityName:string;  // 新内容    
    transDate:string;  // 变更时间    
    createdBy:string;  // 操作时间    
    linkedNurseChangeLogNurse:LinkedNurseChangeLogNurse;    
    linkedNurseChangeLogDepartment:LinkedNurseChangeLogDepartment;    
}


export interface NurseChangeLogInfoData {    
    id:number|undefined;        //    
    nurseId:number|undefined; //    
    linkedNurseId:LinkedNurseChangeLogNurse; //    
    changeType:number|undefined; // 变更类型    
    fromAreaId:number|undefined; //    
    fromDeptId:number|undefined; //    
    linkedFromDeptId:LinkedNurseChangeLogDepartment; //    
    fromEntityId:number|undefined; //    
    fromEntityName:string|undefined; // 原内容    
    toAreaId:number|undefined; //    
    toDeptId:number|undefined; //    
    linkedToDeptId:LinkedNurseChangeLogDepartment; //    
    toEntityId:number|undefined; //    
    toEntityName:string|undefined; // 新内容    
    transDate:string|undefined; // 变更时间    
    status:boolean; //    
    createdAt:string|undefined; //    
    createdBy:string|undefined; // 操作时间    
    updatedAt:string|undefined; //    
    updatedBy:string|undefined; //    
    deletedAt:string|undefined; //    
    deletedBy:string|undefined; //    
    linkedNurseChangeLogNurse:LinkedNurseChangeLogNurse;    
    linkedNurseChangeLogDepartment:LinkedNurseChangeLogDepartment;    
}


export interface LinkedNurseChangeLogNurse {	
    nurseId:number|undefined;    //	
}


export interface LinkedNurseChangeLogDepartment {	
    deptId:number|undefined;    //	
}


export interface NurseChangeLogTableDataState {
    ids:any[];
    tableData: {
        data: Array<NurseChangeLogTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            nurseId: number|undefined;            
            changeType: number|undefined;            
            fromAreaId: number|undefined;            
            fromDeptId: number|undefined;            
            toAreaId: number|undefined;            
            toDeptId: number|undefined;            
            transDate: string|undefined;            
            status: number|undefined;            
            createdAt: string|undefined;            
            createdBy: string|undefined;            
            deletedBy: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface NurseChangeLogEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:NurseChangeLogInfoData;
    rules: object;
}