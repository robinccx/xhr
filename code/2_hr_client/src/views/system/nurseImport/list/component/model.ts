export interface NurseImportTableColumns {    
    nurseId:number;  //    
    sessionId:string;  //    
    nurseCode:string;  //    
    nurseName:string;  //    
    startDate:string;  //    
    endDate:string;  //    
    sex:string;  //    
    birthday:string;  //    
    areaId:number;  //    
    areaCode:string;  //    
    deptId:number;  //    
    deptCode:string;  //    
    educationId:number;  //    
    educationCode:string;  //    
    titleId:number;  //    
    titleCode:string;  //    
    staffType:number;  //    
    staffTypeCode:string;  //    
    note:string;  //    
    status:number;  //    
    createdAt:string;  //    
    createdBy:string;  //    
}


export interface NurseImportInfoData {    
    nurseId:number|undefined;        //    
    sessionId:string|undefined; //    
    nurseCode:string|undefined; //    
    nurseName:string|undefined; //    
    startDate:string|undefined; //    
    endDate:string|undefined; //    
    sex:string|undefined; //    
    birthday:string|undefined; //    
    areaId:number|undefined; //    
    areaCode:string|undefined; //    
    deptId:number|undefined; //    
    deptCode:string|undefined; //    
    educationId:number|undefined; //    
    educationCode:string|undefined; //    
    titleId:number|undefined; //    
    titleCode:string|undefined; //    
    staffType:number|undefined; //    
    staffTypeCode:string|undefined; //    
    note:string|undefined; //    
    status:boolean; //    
    createdAt:string|undefined; //    
    createdBy:string|undefined; //    
}


export interface NurseImportTableDataState {
    nurseIds:any[];
    tableData: {
        data: Array<NurseImportTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            nurseId: number|undefined;            
            sessionId: string|undefined;            
            nurseCode: string|undefined;            
            nurseName: string|undefined;            
            startDate: string|undefined;            
            endDate: string|undefined;            
            sex: string|undefined;            
            birthday: string|undefined;            
            areaId: number|undefined;            
            areaCode: string|undefined;            
            deptId: number|undefined;            
            deptCode: string|undefined;            
            educationId: number|undefined;            
            educationCode: string|undefined;            
            titleId: number|undefined;            
            titleCode: string|undefined;            
            staffType: number|undefined;            
            staffTypeCode: string|undefined;            
            note: string|undefined;            
            status: number|undefined;            
            createdAt: string|undefined;            
            createdBy: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface NurseImportEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:NurseImportInfoData;
    rules: object;
}