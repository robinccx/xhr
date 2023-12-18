export interface EducationTableColumns {    
    educationId:number    
    seq:number;  // 次序    
    educationName:string;  // 学历    
    status:number;  // 状态    
}


export interface EducationInfoData {    
    educationId:number|undefined;        // ID    
    educationName:string|undefined; // 学历    
    seq:number|undefined; // 次序    
    status:boolean; // 状态    
    createdAt:string|undefined; //    
    createdBy:string|undefined; //    
    updatedAt:string|undefined; //    
    updatedBy:string|undefined; //    
    deletedAt:string|undefined; //    
    deletedBy:string|undefined; //    
}


export interface EducationTableDataState {
    educationIds:any[];
    tableData: {
        data: Array<EducationTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            educationName: string|undefined;            
            status: number|undefined;            
            dateRange: string[];
        };
    };
}


export interface EducationEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:EducationInfoData;
    rules: object;
}