/*************************************************************************
	> File Name: manager.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Thu 25 Oct 2012 09:43:40 PM CST
 ************************************************************************/
package library
import "errors"
type MusicManager struct{
	musics []MusicEntry
}
func NewMusicManager() *MusicManager{
	return &MusicManager{make([]MusicEntry,0)}
}
func (p *MusicManager) Len() int{
	return len(p.musics)
}
func (p *MusicManager) Get(index int) (music *MusicEntry,err error){
	if index < 0 || index > p.Len(){
		return nil,errors.New("Index out of range.")
	}
	return &p.musics[index],nil
}

func (p *MusicManager) Find(name string) *MusicEntry{
	if p.Len() == 0{
		return nil
	}
	for _,m := range p.musics{
		if m.Name == name{
			return &m
		}
	}
	return nil
}

func (p *MusicManager) Add(music *MusicEntry){
	p.musics = append(p.musics,*music)
}

func (p *MusicManager) Remove(index int) *MusicEntry{
	if index < 0 || index >= p.Len(){
		return nil
	}
	removedMusic := &p.musics[index]

	//this place , algorithm has a problem, need to be changed
	//slice remove
	if index < p.Len() -1 {
		p.musics = append(p.musics[:index-1],p.musics[index+1:]...)
	}else if index == 0{
		p.musics = make([]MusicEntry,0)
	}else{
		p.musics = p.musics[:index-1]	
	}
	return removedMusic
}
