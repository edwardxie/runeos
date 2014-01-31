function keyPress () {
 if(!(event.keyCode==46)&&!(event.keyCode==8)&&!(event.keyCode==37)&&!(event.keyCode==39))
   if(!((event.keyCode>=48&&event.keyCode<=57)||(event.keyCode>=96&&event.keyCode<=105)))
        event.returnValue=false;
}

/*
<script language="JavaScript" type="text/javascript"> 
<!-- 
function checkMoney(obj){ 
    var tempValue=obj.value.replace(/(^\s+)|(\s+$)/g,''); 
    if(!tempValue){return;} 
    if(/^-?\d+(\.\d+)?$/.test(tempValue)){ 
        obj.value=parseFloat(tempValue).toFixed(2); 
    }else{ 
        alert('请输入合法的货币值！'); 
        obj.select(); 
        return; 
    } 
} 
// --> 
</script> 
触发事件 
onblur="checkMoney(this)" 
*/

/**
 * 根据传入的数值转换成货币显示格式
 * 不能传入的字符形如 “abc  000  123.  .123”

function formatNum(s){
    s=(Math.round(s*100)/100).toString();
    //判断参数是否有值
    if( s == '' || s == null || s == '0'){
        return '0.00';
    }else if(( 0 < s && s < 10 ) || ( -10 < s && s < 0 )){//判断参数是否是个位数，除0外(因为0==''成立)
        if(s.indexOf('.') == -1){
            s = s + '.00';
        }else if(s.indexOf('.') == s.length - 2){
            s = s + '0';
        }
        return s;
    }else{//处理二位数以上的格式
        s += '';
        if(!/^(\+|-)?\d+(\.\d+)?$/.test(s)){ throw (new Error(-1,'lt isn\'t Number ! '));}
        var a = s.match(/^(\+|-)?(\d[^\.]+)(\.\d+)?$/),b = a[2],c = '';
        for(var i = b.length - 3;i > -3;i = i - 3){
            c = ',' + b.substring(i ,i + 3) + c;
        }
        var f = ((a[1]||"") + c.substr(1) + (a[3]||""));
        if(f.indexOf('.') == -1){
            f = f + '.00';
        }else if( f.indexOf('.') == f.length-2){
            f = f + '0';
        }
        return f;
    }
} */