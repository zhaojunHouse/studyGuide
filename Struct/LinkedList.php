<?php

/**
 * Created by PhpStorm.
 * User: zhaojun
 * Date: 2018/12/9
 * Time: 上午11:24
 */
class LinkedListNode {
    public $data;
    public $next;

    public function __construct($data = null) {
        $this->data = $data;
        $this->next = null;
    }
}


class  LinkedList {
    public $head;
    public $length;

    public function __construct($head = null) {
        if ($head == null) {
            $this->head = new LinkedListNode();
        } else {
            $this->head = $head;
        }
        $this->length = 0;
    }

    //length

    public function getLength() {
        return $this->length;
    }

    //insert
    public function insert($data) {
        return $this->insertDataAfter($this->head,$data);
    }

    //delete
    public function delete(LinkedListNode $node){
        if ($node == null){
            return false;
        }

        $preNode = $this->getPreNode($node);
        $preNode->next = $node->next;
        unset($node);

        $this->length--;
        return true;
    }


    //insert data after
    public function insertDataAfter(LinkedListNode $originNode, $data) {
        if ($originNode == null) {
            return false;
        }

        $newNode = new LinkedListNode();
        $newNode->data = $data;

        $newNode->next = $originNode->next;
        $originNode->next = $newNode;

        $this->length++;

        return $newNode;
    }

    //insert data before
    public function insertDataBefore(LinkedListNode $originNode, $data) {
        if ($originNode == null) {
            return false;
        }

        $preNode = $this->getPreNode($originNode);

        return $this->insertDataAfter($preNode, $data);

    }

    //get pre node
    public function getPreNode(LinkedListNode $node) {
        if ($node == null) {
            return false;
        }

        $curNode = $this->head;
        $preNode = $this->head;

        while ($curNode != $node && $curNode != null) {
            $preNode = $curNode;
            $curNode = $curNode->next;
        }

        return $preNode;
    }


    //get node by index
    public function getNodeByIndex($index) {
        if ($index < 0 || $index > $this->length) {
            return false;
        }

        $curNode = $this->head->next;
        for ($i = 0; $i < $index; $i++) {
            $curNode = $curNode->next;
        }

        return $curNode;

    }


    //print
    public function printList(){
        if (null == $this->head->next) {
            return false;
        }

        $curNode = $this->head;
        while ($curNode->next != null) {
            echo $curNode->next->data . ' -> ';

            $curNode = $curNode->next;
        }
        echo 'NULL' . PHP_EOL;

        return true;
    }


    public function findIndex($arr, $len, $val) {
        if (empty($arr) || $len <= 0) {
            return -1;
        }

        if ($arr[$len - 1] == $val) {
            return $len - 1;
        }

        $temp = $arr[$len - 1];
        $arr[$len - 1] = $val;

        $i = 0;
        // O（n）
        while ($arr[$i] != $val) {
            ++$i;
        }

        $arr[$len - 1] = $temp;

        if ($i == $len - 1) {
            return -1;
        } else {
            return $i;
        }

    }


}


$linkedList = new LinkedList();

// test  find index in array
$arr = [
    4,
    2,
    3,
    5,
    9,
    6
];
$a = $linkedList->findIndex($arr, 6, 3);
$b = $linkedList->findIndex($arr, 6, 8);
//var_dump($a);
//var_dump($b);


$linkedList->insert(4);
$linkedList->insert(2);
$linkedList->insert(3);
$linkedList->insert(5);
$linkedList->insert(9);
$linkedList->insert(6);
echo $linkedList->getLength()."\n";
$linkedList->printList();
$node = $linkedList->getNodeByIndex(3);
$linkedList->delete($node);
$linkedList->printList();

