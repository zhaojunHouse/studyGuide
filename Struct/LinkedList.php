<?php
/**
 * Created by PhpStorm.
 * User: zhaojun
 * Date: 2018/12/9
 * Time: 上午11:24
 */

class LinkedListNode
{
    public $data;
    public $next;

    public function __construct($data = null)
    {
        $this->data = $data;
        $this->next = null;
    }
}


class  LinkedList
{
    public $head;
    public $length;

    public function __construct($head = null)
    {
        if ($head == null) {
            $this->head = new LinkedListNode();
        } else {
            $this->head = $head;
        }
        $this->length = 0;
    }

    public function findIndex($arr,$len,$val){
        if(empty($arr) || $len <=0){
            return -1;
        }

        if($arr[$len-1] == $val){
            return $len-1;
        }

        $temp = $arr[$len -1];
        $arr[$len-1] = $val;

        $i = 0;
        while ($arr[$i] != $val){
            ++$i;
        }

        $arr[$len-1] = $temp;

        if ($i == $len-1){
            return -1;
        }else{
            return $i;
        }

    }


}


$linkedList = new LinkedList();
$arr = array(4, 2, 3, 5, 9, 6);
$a = $linkedList->findIndex($arr,6,3);
$b = $linkedList->findIndex($arr,6,8);
var_dump($a);
var_dump($b);
