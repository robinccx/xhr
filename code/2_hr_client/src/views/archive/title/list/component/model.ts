export interface TitleTableColumns {    
    titleId:number    
    seq:number;  // 次序    
    titleName:string;  // 职称    
    status:number;  // 状态    
    createdAt:string;  // 创建时间    
    createdBy:string;  // 创建用户    
}


export interface TitleInfoData {    
    titleId:number|undefined;        // ID    
    titleName:string|undefined; // 职称    
    seq:number|undefined; // 次序    
    status:boolean; // 状态    
    createdAt:string|undefined; // 创建时间    
    createdBy:string|undefined; // 创建用户    
    updatedAt:string|undefined; //    
    updatedBy:string|undefined; // 修改用户    
    deletedAt:string|undefined; //    
    deletedBy:string|undefined; // 删除用户    
}


export interface TitleTableDataState {
    titleIds:any[];
    tableData: {
        data: Array<TitleTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            titleName: string|undefined;            
            status: number|undefined;            
            dateRange: string[];
        };
    };
}


export interface TitleEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:TitleInfoData;
    rules: object;
}