export interface PositionTableColumns {    
    positionId:number    
    positionName:string;  // 职务    
    seq:number;  // 次序    
    status:number;  // 状态    
}


export interface PositionInfoData {    
    positionId:number|undefined;        // 职务ID    
    positionCode:string|undefined; // 职务代码    
    positionName:string|undefined; // 职务    
    seq:number|undefined; // 次序    
    status:boolean; // 状态    
    createdAt:string|undefined; // 创建时间    
    createdBy:string|undefined; // 创建用户    
    updatedAt:string|undefined; //    
    updatedBy:string|undefined; //    
    deletedAt:string|undefined; //    
    deletedBy:string|undefined; //    
}


export interface PositionTableDataState {
    positionIds:any[];
    tableData: {
        data: Array<PositionTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            positionName: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface PositionEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:PositionInfoData;
    rules: object;
}