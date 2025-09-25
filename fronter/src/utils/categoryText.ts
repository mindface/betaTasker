
export function categoryText(text:string):string{
  switch (text) {
   case 'title':
    return 'タスクタイトル'
   case 'text':
    return '情報'
   case 'name':
    return 'カテゴリ'
   case 'disc':
    return '詳細'
   case 'imgPath':
     return 'パス'
   case 'id':
    return 'ID'
   default:
    return 'ID'
  }
}

export const cardCategory = [
  {id:1, value:'todo[タスク]'},
  {id:2, value:'情報の言語化'},
  {id:3, value:'調査情報タスク'},
  {id:4, value:'イメージ構成への関与'},
  {id:5, value:'メタ認知によるタスクのフィードバック'}
]
